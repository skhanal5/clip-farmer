package service

import (
	"encoding/json"
	"github.com/skhanal5/clip-farmer/config"
	"github.com/skhanal5/clip-farmer/internal/api"
	model "github.com/skhanal5/clip-farmer/model/twitch"
	"log"
	"net/http"
)

const (
	twitchClipsAPI = "https://api.twitch.tv/helix/clips"
	twitchUsersAPI = "https://api.twitch.tv/helix/users"
)

type TwitchService struct {
	headers      map[string]string
	twitchClient *api.TwitchClient
}

func NewTwitchService(config config.Config) *TwitchService {
	log.Print("Building new TwitchService instance")
	service := &TwitchService{
		headers:      make(map[string]string),
		twitchClient: api.NewTwitchClient(),
	}
	service.headers["Authorization"] = "Bearer " + config.TwitchOAuthConfig.AccessToken
	service.headers["Client-Id"] = config.TwitchClientId
	return service
}

func (service *TwitchService) FetchUsers(user string) (model.TwitchGraphQLResponse, error) {
	log.Print("Fetching user details for: " + user)

	req := service.buildUsersRequest(user)
	body, err := service.twitchClient.SendGetRequest(req)
	if err != nil {
		return model.TwitchGraphQLResponse{}, err
	}

	var gqlResponse model.TwitchGraphQLResponse
	err = json.Unmarshal(body, &gqlResponse)
	if err != nil {
		return model.TwitchGraphQLResponse{}, err
	}
	return gqlResponse, nil
}

func (service *TwitchService) FetchUserClips(broadcasterId string) (model.TwitchGraphQLResponse, error) {
	log.Print("Fetching clips from user with broadcastId: " + broadcasterId)
	clipsEndpoint := service.buildClipsRequest(broadcasterId)
	responseBody, err := service.twitchClient.SendGetRequest(clipsEndpoint)
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

func (service *TwitchService) buildUsersRequest(user string) *http.Request {
	request, _ := http.NewRequest("GET", twitchUsersAPI, nil)

	queryParams := request.URL.Query()
	queryParams.Add("login", user)
	request.URL.RawQuery = queryParams.Encode()

	setRequestHeaders(request, service.headers)
	return request
}

func (service *TwitchService) buildClipsRequest(broadcasterId string) *http.Request {
	request, _ := http.NewRequest("GET", twitchClipsAPI, nil)

	queryParams := request.URL.Query()
	queryParams.Add("broadcaster_id", broadcasterId)
	request.URL.RawQuery = queryParams.Encode()

	setRequestHeaders(request, service.headers)
	return request
}

func setRequestHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Add(key, value)
	}
}
