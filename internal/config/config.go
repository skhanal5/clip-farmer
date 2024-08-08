package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	TikTokClientKey     string
	TikTokClientSecret  string
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
		viper.GetString("secrets.twitch.client-id"),
		viper.GetString("secrets.twitch.client-oauth"),
		viper.GetString("query.twitch-creator"),
	}
}
