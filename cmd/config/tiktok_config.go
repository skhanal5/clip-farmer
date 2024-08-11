package config

import (
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	clientKey    string
	clientSecret string
)

// tiktokCmd represents the tiktok subcommand
var tiktokConfigCmd = &cobra.Command{
	Use:   "tiktok",
	Short: "Configure TikTok environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		clientKey, _ := cmd.Flags().GetString("client-key")
		clientSecret, _ := cmd.Flags().GetString("client-secret")
		if clientKey != "" {
			viper.Set("secrets.tiktok.client-key", clientKey)
		}
		if clientSecret != "" {
			viper.Set("secrets.tiktok.client-secret", clientSecret)
		}
		SaveConfig("./config.yaml")
	},
}

func init() {
	configCmd.AddCommand(tiktokConfigCmd)
	tiktokConfigCmd.Flags().StringVarP(&clientKey, "client-key", "k", "",
		"Set the client-key of the TikTok app that we want to connect to.")
	tiktokConfigCmd.Flags().StringVarP(&clientSecret, "client-secret", "s", "",
		"Set the client-secret of the TikTok app that we want to connect to.")
	tiktokConfigCmd.MarkFlagsRequiredTogether("client-key", "client-secret")
}
