package telegram

import (
        "log"
        "time"

        "github.com/yanzay/tbot"
        "github.com/luanguimaraesla/memoir/model"
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

func questionsHandler(m *tbot.Message){
        for i, q := range talk.Questions {
                time.Sleep(time.Duration(1) * time.Second)
                m.Replyf("%d. %s", i, q.Text)
        }
        m.Reply("This is all I have. :)")
}

func Run(t *model.Talk, token string){
        talk = t
        runTelegramServer(token)
}
