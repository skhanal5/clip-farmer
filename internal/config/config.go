package config

import (
	"log"
	"os"
)

type Config struct {
	TiktokClientKey    string
	TikTokClientSecret string
	TwitchClientId     string
	TwitchClientOAuth  string
}

func NewConfig() Config {
	log.Print("Loading environment variable")
	return Config{
		os.Getenv("TIKTOK_CLIENT_KEY"),
		os.Getenv("TIKTOK_CLIENT_SECRET"),
		os.Getenv("TWITCH_CLIENT_ID"),
		os.Getenv("TWITCH_CLIENT_OAUTH"),
	}
}
