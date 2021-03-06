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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
        "github.com/luanguimaraesla/memoir/telegram"
        "github.com/luanguimaraesla/memoir/model"
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
        var talk model.Talk
        err := viper.Unmarshal(&talk)
        if err != nil {
                log.Panic("unable to unmarshal config")
        }

        exporterAddr, err := cmd.Flags().GetString("exporter")
        if err != nil {
                log.Panic("you should configure an metrics exporter (--exporter)")
        }

        switch agent, _ := cmd.Flags().GetString("agent"); agent {
        case "telegram":
                token, _ := cmd.Flags().GetString("token")
                telegram.Run(&talk, token, exporterAddr)
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
        runCmd.Flags().StringP("exporter", "e", "", "Exporter address")
}
