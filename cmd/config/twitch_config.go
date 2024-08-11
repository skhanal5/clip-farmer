package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	clientId    string
	clientOAuth string
)

// twitchCmd represents the twitch subcommand
var twitchConfigCmd = &cobra.Command{
	Use:   "twitch",
	Short: "Configure Twitch environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		clientId, _ := cmd.Flags().GetString("client-id")
		clientOAuth, _ := cmd.Flags().GetString("client-oauth")
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
	twitchConfigCmd.Flags().StringVarP(&clientId, "client-id", "i", "",
		"Set the client-id of the Twitch user that we want to send authenticated requests from.")
	twitchConfigCmd.Flags().StringVarP(&clientOAuth, "client-oauth", "o", "",
		"Set the client-oauth of the Twitch user that we want to send authenticated requests from.")
}
