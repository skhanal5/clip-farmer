package config

import (
	"errors"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		if clientId == "" || clientOAuth == "" {
			return errors.New("you must provide a client id and client oauth")
		}
		viper.Set("secrets.twitch.client-id", clientId)
		viper.Set("secrets.twitch.client-oauth", clientOAuth)
		SaveConfig("./config.yaml")
		return nil
	},
}

func init() {
	configCmd.AddCommand(twitchConfigCmd)
	twitchConfigCmd.Flags().StringVarP(&clientId, "client-id", "i", "",
		"Set the client-id of the Twitch user that we want to send authenticated requests from.")
	twitchConfigCmd.Flags().StringVarP(&clientOAuth, "client-oauth", "o", "",
		"Set the client-oauth of the Twitch user that we want to send authenticated requests from.")
	twitchConfigCmd.MarkFlagsRequiredTogether("client-id", "client-oauth")
}
