package fetch

import (
	"errors"
	"github.com/skhanal5/clip-farmer/manager"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	user string
)

// twitchCmd represents the twitch subcommand
var twitchCmd = &cobra.Command{
	Use:   "twitch",
	Short: "Set TikTok configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		manager, err := buildManager()

		if err != nil {
			return err
		}

		if user != "" {
			manager.FetchAndDownloadClips(user)
		}
		return nil
	},
}

func init() {
	fetchCmd.AddCommand(twitchCmd)
	twitchCmd.Flags().StringVarP(&user, "user", "u", "", "Twitch username")
}

func buildManager() (manager.TwitchManager, error) {
	clientId := viper.GetString("secrets.twitch.client-id")
	clientOAuth := viper.GetString("secrets.twitch.client-oauth")

	if clientId == "" {
		return manager.TwitchManager{}, errors.New("twitch clientId is not configured")
	}

	if clientOAuth == "" {
		return manager.TwitchManager{}, errors.New("twitch clientOAuth is not configured")
	}

	return manager.InitTwitchManager(clientId, clientOAuth), nil
}
