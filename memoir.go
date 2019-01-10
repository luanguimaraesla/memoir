package main

import (
        "log"
        "os"
        "time"
        "strconv"
        "github.com/yanzay/tbot"
)

func timeHandler(m *tbot.Message){
        // m.Vars contains all variables, parsed during routing
        secondsStr := m.Vars["seconds"]
        // convert string variable to integer seconds value
        seconds, err := strconv.Atoi(secondsStr)
        if err != nil {
                m.Reply("Invalid number of seconds")
                return
        }
        m.Replyf("Timer for %d seconds started", seconds)
        time.Sleep(time.Duration(seconds) * time.Second)
        m.Reply("Timeout")
}

func main(){
        bot, err := tbot.NewServer(os.Getenv("TELEGRAM_TOKEN"))
        if err != nil {
                log.Fatal(err)
        }
        bot.Handle("/answer", "42")
        bot.HandleFunc("/timer {seconds}", timeHandler)
        bot.ListenAndServe()
}
