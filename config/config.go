package config

import (
	model "github.com/skhanal5/clip-farmer/model/twitch"
	"log"
	"os"
)

type Config struct {
	TiktokApiKey       string
	TwitchClientId     string
	TwitchClientSecret string
	TwitchOAuthConfig  model.TwitchOAuthResponse
}

func LoadConfig() Config {
	log.Print("Loading environment variable")
	tiktokAPIKey := os.Getenv("TIKTOK_API_KEY")
	twitchClientId := os.Getenv("TWITCH_CLIENT_ID")
	twitchClientSecret := os.Getenv("TWITCH_CLIENT_SECRET")
	return Config{
		tiktokAPIKey,
		twitchClientId,
		twitchClientSecret,
		model.TwitchOAuthResponse{},
	}
}

func (config *Config) SetTwitchBearerToken(oauthResponse model.TwitchOAuthResponse) {
	config.TwitchOAuthConfig = oauthResponse
}
