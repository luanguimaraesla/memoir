package telegram

import (
        "log"
        "fmt"
        "sync"
        "strconv"

        "github.com/yanzay/tbot"
        "github.com/luanguimaraesla/memoir/model"
        "github.com/luanguimaraesla/memoir/collectorclient"
)


type talkController struct {
        talk *model.Talk
        responseValue float32
        callbackMutex *sync.Mutex
}

var gtc *talkController

func runTelegramServer(token string){
        bot, err := tbot.NewServer(token)
        if err != nil {
                log.Fatal(err)
        }
        bot.Handle("answer", "42")
        bot.HandleFunc("ask", questionsHandler)
        bot.HandleFunc("/{value}", callbackValueHandler)
        bot.ListenAndServe()
}

func callbackValueHandler(m *tbot.Message) {
        value, err := strconv.ParseFloat(m.Vars["value"], 32)
        if err != nil {
                log.Panicf("can't convert %s to float32", m.Vars["value"])
                gtc.responseValue = 0.0
        } else {
                gtc.responseValue = float32(value)
        }
        gtc.callbackMutex.Unlock()
}

func ask(m *tbot.Message) chan *model.Measure {
        yield := make(chan *model.Measure)
        go func() {
                for i, q := range gtc.talk.Questions {
                        gtc.callbackMutex.Lock()
                        text := fmt.Sprintf("%d. %s", i, q.Text)

                        buttons := [][]string{
                                {"/1", "/2", "/3"},
                                {"/4", "/5", "/6"},
                        }
                        m.ReplyKeyboard(text, buttons, tbot.OneTimeKeyboard)
                        gtc.callbackMutex.Lock()
                        m.Replyf("Saved value: %f", gtc.responseValue)
                        yield <- &model.Measure{
                               Reference: &q,
                               Value: gtc.responseValue,
                        }
                        gtc.callbackMutex.Unlock()
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
        gtc = &talkController{
                talk: t,
                callbackMutex: &sync.Mutex{},
        }
        runTelegramServer(token)
}
