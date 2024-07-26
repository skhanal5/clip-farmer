package service

import (
	"encoding/json"
	"fmt"
	"github.com/skhanal5/clip-farmer/config"
	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/request"
	model "github.com/skhanal5/clip-farmer/model/twitch"
	"log"
)

const (
	twitchClipsAPI      = "https://api.twitch.tv/helix/clips"
	twitchUsersAPI      = "https://api.twitch.tv/helix/users"
	twitchOAuthEndpoint = "https://id.twitch.tv/oauth2/token"
)

const (
	GET  = "GET"
	POST = "POST"
)

type TwitchService struct {
	twitchClient *client.TwitchClient
}

func NewTwitchService() *TwitchService {
	log.Print("Building new TwitchService instance")
	service := &TwitchService{
		twitchClient: client.NewTwitchClient(),
	}
	return service
}

func (service *TwitchService) FetchOAuth(config config.Config) model.TwitchOAuthResponse {
	oauthRequest := buildOAuthRequest(config)
	log.Print("Invoking Twitch OAuth request")
	responseBody, err := service.twitchClient.SendRequest(oauthRequest)
	if err != nil {
		panic(err)
	}
	var oauthResponse model.TwitchOAuthResponse
	err = json.Unmarshal(responseBody, &oauthResponse)
	if err != nil {
		panic(err)
	}
	return oauthResponse
}

func buildOAuthRequest(config config.Config) request.TwitchRequest {
	queryParams := make(map[string]string)
	queryParams["client_id"] = config.TwitchClientId
	queryParams["client_secret"] = config.TwitchClientSecret
	queryParams["grant_type"] = "client_credentials"
	oauthRequest := request.TwitchRequestData{
		RequestURL:  twitchOAuthEndpoint,
		RequestType: POST,
		Query:       queryParams,
	}
	return oauthRequest.BuildRequest()
}

func (service *TwitchService) FetchUsers(user string, config config.Config) model.TwitchUserResponse {
	userRequest := buildGetRequest(map[string]string{"login": user}, twitchUsersAPI, config)
	fmt.Println(userRequest.Request)
	log.Print("Invoking GET Users")
	body, err := service.twitchClient.SendRequest(userRequest)
	if err != nil {
		panic(err)
	}

	var gqlResponse model.TwitchUserResponse
	err = json.Unmarshal(body, &gqlResponse)
	if err != nil {
		panic(err)
	}

	return gqlResponse
}

func (service *TwitchService) FetchUserClips(broadcasterId string, config config.Config) model.TwitchClipResponse {
	clipsRequest := buildGetRequest(map[string]string{"broadcaster_id": broadcasterId}, twitchClipsAPI, config)

	log.Print("Invoking GET Clips")
	responseBody, err := service.twitchClient.SendRequest(clipsRequest)
	if err != nil {
		panic(err)
	}

	var gqlResponse model.TwitchClipResponse
	err = json.Unmarshal(responseBody, &gqlResponse)

	if err != nil {
		panic(err)
	}
	return gqlResponse
}

func buildAuthorizationHeaders(config config.Config) map[string][]string {
	headers := make(map[string][]string)
	headers["Authorization"] = []string{"Bearer " + config.TwitchOAuthConfig.AccessToken}
	headers["Client-Id"] = []string{config.TwitchClientId}
	return headers
}

func buildGetRequest(queryParams map[string]string, url string, config config.Config) request.TwitchRequest {
	headers := buildAuthorizationHeaders(config)
	requestData := request.TwitchRequestData{
		RequestURL:  url,
		RequestType: GET,
		Query:       queryParams,
		Headers:     headers,
	}
	return requestData.BuildRequest()
}
