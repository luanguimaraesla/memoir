// Copyright © 2019 Luan Guimarães Lacerda
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
        "log"
        "os"
        "time"
        "strconv"
        "github.com/yanzay/tbot"
        "github.com/luanguimaraesla/memoir/memoir/cmd"
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
	cmd.Execute()
}
