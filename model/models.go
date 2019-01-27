package model

type Question struct {
        Text string
        Kind string
        Group string
        Metric string
        Repeat string
}

type Talk struct {
        Questions []Question
}

type Measure struct {
        Reference *Question
        Value float32
}
