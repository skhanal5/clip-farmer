package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	user        string
	clientId    string
	clientOAuth string
)

// twitchCmd represents the twitch command
var twitchConfigCmd = &cobra.Command{
	Use:   "twitch",
	Short: "Set Twitch configuration",
	Run: func(cmd *cobra.Command, args []string) {
		creatorFlag, _ := cmd.Flags().GetString("creator")
		clientId, _ := cmd.Flags().GetString("client-id")
		clientOAuth, _ := cmd.Flags().GetString("client-oauth")
		if creatorFlag != "" {
			viper.Set("secrets.twitch.creator", creatorFlag)
		}
		if clientId != "" {
			viper.Set("secrets.twitch.client-id", clientId)
		}
		if clientOAuth != "" {
			viper.Set("secrets.twitch.client-oauth", clientOAuth)
		}
		SaveConfig("./config.yaml")
	},
}

func init() {
	configCmd.AddCommand(twitchConfigCmd)
	twitchConfigCmd.Flags().StringVarP(&user, "creator", "c", "",
		"Set the username of the Twitch creator that we want to fetch clips from.")
	twitchConfigCmd.Flags().StringVarP(&clientId, "client-id", "i", "",
		"Set the client-id of the Twitch user that we want to send authenticated requests from.")
	twitchConfigCmd.Flags().StringVarP(&clientOAuth, "client-oauth", "o", "",
		"Set the client-oauth of the Twitch user that we want to send authenticated requests from.")
}
