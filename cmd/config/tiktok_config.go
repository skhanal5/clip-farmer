package config

import (
	"errors"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		if clientKey == "" || clientSecret == "" {
			return errors.New("you must provide a client key and client secret")
		}
		viper.Set("secrets.tiktok.client-key", clientKey)
		viper.Set("secrets.tiktok.client-secret", clientSecret)
		SaveConfig("./config.yaml")
		return nil
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
