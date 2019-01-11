// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
        "log"
        "time"
        "strconv"

	"github.com/spf13/cobra"
        "github.com/yanzay/tbot"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run server memoir bot",
	Long: `Memoir server will run the bot to listen for your direct
messages on Telegram, RocketChat, Discord or Slack.`,
	Run: func(cmd *cobra.Command, args []string) {
                runServer(cmd, args)
	},
}

func runServer(cmd *cobra.Command, args []string){
        switch agent, _ := cmd.Flags().GetString("agent"); agent {
        case "telegram":
                token, _ := cmd.Flags().GetString("token")
                runTelegramServer(token)
        case "rocketchat":
                log.Fatal("Not implemented yet: %s", agent)
        case "slack":
                log.Fatal("Not implemented yet: %s", agent)
        case "discord":
                log.Fatal("Not implemented yet: %s", agent)
        default:
                log.Fatal("Invalid agent: %s", agent)
        }
}

func runTelegramServer(token string){
        bot, err := tbot.NewServer(token)
        if err != nil {
                log.Fatal(err)
        }
        bot.Handle("/answer", "42")
        bot.HandleFunc("/timer {seconds}", timeHandler)
        bot.ListenAndServe()
}

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

func init() {
	rootCmd.AddCommand(runCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
        runCmd.Flags().StringP("agent", "a", "telegram", "Chat tool you'll talk to bot")
        runCmd.Flags().StringP("token", "t", "", "Chat token")
}
