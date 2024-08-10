package cmd

import (
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	clientKey    string
	clientSecret string
)

// tiktokCmd represents the tiktok command
var tiktokConfigCmd = &cobra.Command{
	Use:   "tiktok",
	Short: "Set TikTok configuration",
	Run: func(cmd *cobra.Command, args []string) {
		clientKey, _ := cmd.Flags().GetString("client-id")
		clientSecret, _ := cmd.Flags().GetString("client-oauth")
		if clientKey != "" {
			viper.BindEnv("secrets.twitch.client-id", clientId)
		}
		if clientSecret != "" {
			viper.BindEnv("secrets.twitch.client-oauth", clientOAuth)
		}
		SaveConfig()
	},
}

func init() {
	tiktokConfigCmd.Flags().StringVarP(&clientKey, "client-key", "k", "", "Set the client-key of the TikTok app that we want to connect to.")
	tiktokConfigCmd.Flags().StringVarP(&clientSecret, "client-secret", "s", "", "Set the client-secret of the TikTok app that we want to connect to.")
	configCmd.AddCommand(tiktokConfigCmd)
}
