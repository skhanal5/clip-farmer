package config

import (
	"github.com/skhanal5/clip-farmer/internal/twitchmodel"
	"log"
	"os"
)

type Config struct {
	TiktokClientKey    string
	TikTokClientSecret string
	TwitchClientId     string
	TwitchClientSecret string
	TwitchOAuthConfig  twitchmodel.TwitchOAuthResponse
}

func NewConfig() Config {
	log.Print("Loading environment variable")
	return Config{
		os.Getenv("TIKTOK_CLIENT_KEY"),
		os.Getenv("TIKTOK_CLIENT_SECRET"),
		os.Getenv("TWITCH_CLIENT_ID"),
		os.Getenv("TWITCH_CLIENT_SECRET"),
		twitchmodel.TwitchOAuthResponse{},
	}
}

func (config *Config) SetTwitchBearerToken(oauthResponse twitchmodel.TwitchOAuthResponse) {
	config.TwitchOAuthConfig = oauthResponse
}
