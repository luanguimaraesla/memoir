package telegram

import (
        "log"
        "time"
        "strconv"

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
        bot.HandleFunc("ask {seconds}", questionsHandler)
        bot.ListenAndServe()
}

func questionsHandler(m *tbot.Message){
        // m.Vars contains all variables, parsed during routing
        secondsStr := m.Vars["seconds"]
        // convert string variable to integer seconds value
        seconds, err := strconv.Atoi(secondsStr)
        if err != nil {
                m.Reply("Invalid number of seconds")
                return
        }

        for i, q := range talk.Questions {
                m.Replyf("%d. %s", i, q.Text)
                time.Sleep(time.Duration(seconds) * time.Second)
        }
        m.Reply("This is all I have. :)")
}

func Run(t *model.Talk, token string){
        talk = t
        runTelegramServer(token)
}
