package collectorclient

import (
        "log"
        "fmt"
        "context"
        "strings"
        "time"

        "google.golang.org/grpc"

        pb "github.com/luanguimaraesla/memoir/metricsgateway"
        "github.com/luanguimaraesla/memoir/model"
)

type nextMeasureFunc func() chan *model.Measure

type collectorClient struct {
        Addr string
        cf nextMeasureFunc
        cc pb.MetricsGatewayClient
        conn *grpc.ClientConn
}

type CollectorClient interface {
        SendMeasures()
        Close()
}

var mmt = map[string]pb.Measure_MetricType{
        "gauge": pb.Measure_GAUGE,
        "counter": pb.Measure_COUNTER,
        "histogram": pb.Measure_HISTOGRAM,
        "summary": pb.Measure_SUMMARY,
}

func buildArgs(m *model.Measure) *pb.Measure {
        return &pb.Measure{
                Name: fmt.Sprintf("%s.%s", m.Reference.Group, m.Reference.Metric),
                Kind: mmt[strings.ToLower(m.Reference.Kind)],
                Value: m.Value,
        }
}

func (c *collectorClient) SendMeasures() {
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        defer cancel()

        stream, err := c.cc.AddMeasure(ctx)
        if err != nil {
                log.Fatalf("%v.AddMeasure(_) = _, %v", c.cc, err)
        }

        for m := range c.cf() {
                a := buildArgs(m)
                if err := stream.Send(a); err != nil {
                        log.Fatalf("%v.Send(%v) = %v", stream, a, err)
                }
        }

        reply, err := stream.CloseAndRecv()
        if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v", reply)
}

func (c *collectorClient) Close() {
        c.conn.Close()
}

func NewCollectorClient(addr string, collectorFunc nextMeasureFunc) CollectorClient{
        log.Printf("connecting to remote metrics gateway on %s", addr)
        conn, err := grpc.Dial(addr, grpc.WithInsecure())
        if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

        return &collectorClient{
                Addr: addr,
                cc: pb.NewMetricsGatewayClient(conn),
                cf: collectorFunc,
                conn: conn,
        }
}
