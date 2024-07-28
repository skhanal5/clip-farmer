package config

import (
	model "github.com/skhanal5/clip-farmer/model/twitch"
	"log"
	"os"
)

type Config struct {
	TiktokClientKey    string
	TikTokClientSecret string
	TwitchClientId     string
	TwitchClientSecret string
	TwitchOAuthConfig  model.TwitchOAuthResponse
}

func NewConfig() Config {
	log.Print("Loading environment variable")
	return Config{
		os.Getenv("TIKTOK_CLIENT_KEY"),
		os.Getenv("TIKTOK_CLIENT_SECRET"),
		os.Getenv("TWITCH_CLIENT_ID"),
		os.Getenv("TWITCH_CLIENT_SECRET"),
		model.TwitchOAuthResponse{},
	}
}

func (config *Config) SetTwitchBearerToken(oauthResponse model.TwitchOAuthResponse) {
	config.TwitchOAuthConfig = oauthResponse
}
