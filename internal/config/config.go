package config

import (
	model "github.com/skhanal5/clip-farmer/internal/model/tiktok"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	TiktokClientKey     string
	TikTokClientSecret  string
	TikTokOAuth         model.TikTokOAuthResponse
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
		model.TikTokOAuthResponse{},
		viper.GetString("secrets.twitch.client-id"),
		viper.GetString("secrets.twitch.client-oauth"),
		viper.GetString("query.twitch-creator"),
	}
}
func (c *Config) SetTikTokOAuth(oauth model.TikTokOAuthResponse) {
	c.TikTokOAuth = oauth
}
