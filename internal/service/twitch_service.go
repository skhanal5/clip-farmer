package service

import (
	"github.com/skhanal5/clip-farmer/config"
	"github.com/skhanal5/clip-farmer/internal/api"
	"github.com/skhanal5/clip-farmer/model/twitch"
)

func FetchUser(username string, config config.Config) (model.TwitchUser, error) {
	clips, err := api.GetUser(username, config)
	if err != nil {
		return model.TwitchUser{}, err
	}
	return clips, nil
}

func FetchUserClips(broadcasterId string, config config.Config) (model.TwitchClip, error) {
	clips, err := api.GetTwitchClips(broadcasterId, config)
	if err != nil {
		return model.TwitchClip{}, err
	}
	return clips, nil
}
