package config

import (
	"encoding/json"
	"github.com/skhanal5/clip-farmer/internal/tiktok"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	TikTokClientKey     string
	TikTokClientSecret  string
	TikTokOAuth         tiktok.OAuthResponse
	TwitchClientId      string
	TwitchClientOAuth   string
	TwitchTargetCreator string
}

func NewConfig() Config {
	log.Print("Loading environment variable")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	return Config{
		viper.GetString("secrets.tiktok.client-key"),
		viper.GetString("secrets.tiktok.client-secret"),
		LoadOAuth(),
		viper.GetString("secrets.twitch.client-id"),
		viper.GetString("secrets.twitch.client-oauth"),
		viper.GetString("query.twitch-creator"),
	}
}

func (c *Config) SetTikTokOAuth(oauth tiktok.OAuthResponse) {
	c.TikTokOAuth = oauth
}

func LoadOAuth() tiktok.OAuthResponse {
	var oauthResponse tiktok.OAuthResponse
	file, err := os.Open("tiktok_oauth_resp.json")
	if err != nil {
		log.Print(err)
		return oauthResponse
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&oauthResponse)
	if err != nil {
		log.Print(err)
		return oauthResponse
	}

	// Check if the token is expired
	// Get Refresh token
	return oauthResponse
}
