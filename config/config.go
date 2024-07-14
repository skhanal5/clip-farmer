package config

import (
	"log"
	"os"
)

type Config struct {
	TiktokApiKey      string
	TwitchBearerToken string
	TwitchClientId    string
}

func LoadConfig() Config {
	log.Print("Loading environment variables...")
	tiktokAPIKey := os.Getenv("TIKTOK_API_KEY")
	twitchBearerToken := os.Getenv("TWITCH_BEARER_TOKEN")
	twitchClientId := os.Getenv("TWITCH_CLIENT_ID")
	return Config{
		tiktokAPIKey,
		twitchBearerToken,
		twitchClientId,
	}
}
