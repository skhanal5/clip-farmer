package config

import (
	"log"
	"os"
)

type Config struct {
	TiktokApiKey       string
	TwitchClientId     string
	TwitchClientSecret string
	TwitchBearerToken  string
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
		"",
	}
}

// mutable, probably is an antipattern -> fix this
func (config *Config) SetTwitchClientSecret(secret string) {
	config.TwitchClientSecret = secret
}
