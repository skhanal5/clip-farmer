package service

import (
	"encoding/json"
	"github.com/skhanal5/clip-farmer/config"
	"github.com/skhanal5/clip-farmer/internal/api"
	model "github.com/skhanal5/clip-farmer/model/twitch"
	"log"
)

const (
	twitchClipsAPI = "https://api.twitch.tv/helix/clips"
	twitchUsersAPI = "https://api.twitch.tv/helix/users"
)

type TwitchService struct {
	twitchClient *api.TwitchClient
}

func NewTwitchService(config config.Config) *TwitchService {
	log.Print("Building new TwitchService instance")
	return &TwitchService{
		twitchClient: api.NewTwitchClient(config),
	}
}

func (s *TwitchService) FetchUser(username string) (model.TwitchGraphQLResponse, error) {
	log.Print("Fetching user details for: " + username)
	usersEndpoint := formatUsersEndpoint(username)
	resp, err := s.twitchClient.SendGetRequest(usersEndpoint)
	if err != nil {
		return model.TwitchGraphQLResponse{}, err
	}
	var gqlResponse model.TwitchGraphQLResponse
	err = json.Unmarshal(resp, &gqlResponse)
	if err != nil {
		return model.TwitchGraphQLResponse{}, err
	}
	return gqlResponse, nil
}

func (s *TwitchService) FetchUserClips(broadcasterId string) (model.TwitchGraphQLResponse, error) {
	log.Print("Fetching clips from user with broadcastId: " + broadcasterId)
	clipsEndpoint := formatClipsEndpoint(broadcasterId)
	responseBody, err := s.twitchClient.SendGetRequest(clipsEndpoint)
	if err != nil {
		return model.TwitchGraphQLResponse{}, err
	}
	var gqlResponse model.TwitchGraphQLResponse
	err = json.Unmarshal(responseBody, &gqlResponse)
	if err != nil {
		return model.TwitchGraphQLResponse{}, err
	}
	return gqlResponse, nil
}

func formatClipsEndpoint(broadcasterId string) string {
	return twitchClipsAPI + "?broadcaster_id=" + broadcasterId
}

func formatUsersEndpoint(loginName string) string {
	return twitchUsersAPI + "?login_name=" + loginName
}
