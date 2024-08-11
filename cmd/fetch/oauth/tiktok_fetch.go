package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/skhanal5/clip-farmer/cmd/config"
	"github.com/skhanal5/clip-farmer/internal/tiktok"
	"github.com/skhanal5/clip-farmer/manager"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// tiktokCmd represents the tiktok command
var tiktokCmd = &cobra.Command{
	Use:   "tiktok",
	Short: "Fetch OAuth details from TikTok's API",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := loadExistingPreviousToken()
		if err != nil {
			if err = validateTikTokSecrets(); err != nil {
				return err
			}
			setOAuth()
		}
		config.SaveConfig("./config.yaml")
		return nil
	},
}

func init() {
	oauthCmd.AddCommand(tiktokCmd)
}

func validateTikTokSecrets() error {
	clientKey := viper.IsSet("secrets.tiktok.client-key")
	clientSecret := viper.IsSet("secrets.tiktok.client-secret")

	if !clientKey {
		return errors.New("twitch clientId is not configured")
	}

	if !clientSecret {
		return errors.New("twitch clientOAuth is not configured")
	}
	return nil
}

func setOAuth() {
	clientKey := viper.GetString("secrets.tiktok.client-key")
	clientSecret := viper.GetString("secrets.tiktok.client-secret")
	oauth := manager.FetchTiktokOAuth(clientKey, clientSecret)
	viper.Set("secrets.tiktok.client-oauth", oauth)
}

func loadExistingPreviousToken() error {
	var oauthResponse tiktok.OAuthToken

	file, err := os.Open("tiktok_oauth_resp.json")
	if err != nil {
		return errors.New("cannot find previously cached oauth response")
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&oauthResponse)
	if err != nil {
		return errors.New("failed to deserialize Tiktok OAuth Token")
	}

	// add autorefresh validation logic
	viper.Set("secrets.tiktok.client-oauth", oauthResponse.AccessToken)
	return nil
}

func saveConfig(path string) {
	viper.SetConfigFile(path)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("Failed to write config:", err)
	}
}
