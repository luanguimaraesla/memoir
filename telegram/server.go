package telegram

import (
        "log"
        "time"

        "github.com/yanzay/tbot"
        "github.com/luanguimaraesla/memoir/model"
        "github.com/luanguimaraesla/memoir/collectorclient"
)

var talk *model.Talk

func runTelegramServer(token string){
        bot, err := tbot.NewServer(token)
        if err != nil {
                log.Fatal(err)
        }
        bot.Handle("answer", "42")
        bot.HandleFunc("ask", questionsHandler)
        bot.ListenAndServe()
}

func ask(m *tbot.Message) chan *model.Measure {
        yield := make(chan *model.Measure)
        go func() {
                for i, q := range talk.Questions {
                        time.Sleep(time.Duration(1) * time.Second)
                        m.Replyf("%d. %s", i, q.Text)
                        yield <- &model.Measure{
                               Reference: &q,
                               Value: 1.0,
                        }
                }
                close(yield)
        }()
        return yield
}

func askMascarade (m *tbot.Message) func() chan *model.Measure {
        return func() chan *model.Measure {
                return ask(m)
        }
}

func questionsHandler(m *tbot.Message){
        c := collectorclient.NewCollectorClient("localhost:50051", askMascarade(m))
        defer c.Close()
        c.SendMeasures()
        m.Reply("This is all I have. :)")
}

func Run(t *model.Talk, token string){
        talk = t
        runTelegramServer(token)
}
