package clips

import (
	"errors"

	"github.com/skhanal5/clip-farmer/internal/twitch"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	user string
	period string
	sort string
)

// twitchCmd represents the twitch subcommand
var twitchCmd = &cobra.Command{
	Use:   "twitch",
	Short: "Fetch clips from Twitch",
	RunE: func(cmd *cobra.Command, args []string) error {
		manager, err := buildManager()

		if err != nil {
			return err
		}

		if user != "" {
			manager.FetchAndDownloadClips(user, period, sort)
		} else {
			return errors.New("did not specify a Twitch user")
		}
		return nil
	},
}

func init() {
	clipsCmd.AddCommand(twitchCmd)
	twitchCmd.Flags().StringVarP(&user, "user", "u", "",
		"Twitch username of the creator whose content we want to fetch and download.")
	twitchCmd.Flags().StringVarP(&period, "period", "p", "ALL_TIME",
		"Pass in a filter to get the clips from a time period. Accepted values are LAST_DAY, LAST_WEEK, LAST_MONTH, ALL_TIME.")
	twitchCmd.Flags().StringVarP(&sort, "sort", "s", "TRENDING",
		"Pass in a filter to sort the clips. Accepted values are CREATED_AT_ASC, CREATED_AT_DESC,, VIEWS_ASC, VIEWS_DESC, TRENDING.")
	twitchCmd.MarkFlagsRequiredTogether("period", "sort")
}

func buildManager() (twitch.TwitchManager, error) {
	clientId := viper.GetString("secrets.twitch.client-id")
	clientOAuth := viper.GetString("secrets.twitch.client-oauth")

	if clientId == "" {
		return twitch.TwitchManager{}, errors.New("twitch clientId is not configured")
	}

	if clientOAuth == "" {
		return twitch.TwitchManager{}, errors.New("twitch clientOAuth is not configured")
	}

	return twitch.InitTwitchManager(clientId, clientOAuth), nil
}
