package twitchrequest

import (
	"github.com/skhanal5/clip-farmer/internal/config"
	"github.com/skhanal5/clip-farmer/internal/datamodel"
	"net/http"
)

const (
	twitchClipsAPI      = "https://api.twitch.tv/helix/clips"
	twitchUsersAPI      = "https://api.twitch.tv/helix/users"
	twitchOAuthEndpoint = "https://id.twitch.tv/oauth2/token"
)

func BuildTwitchOAuthRequest(config config.Config) *http.Request {
	data := datamodel.RequestData{
		RequestType:     datamodel.POST,
		RequestURL:      twitchOAuthEndpoint,
		QueryParameters: map[string]string{"client_id": config.TwitchClientId, "client_secret": config.TwitchClientSecret, "grant_type": "client_credentials"},
		Headers:         twitchAuthorizationHeaders(config),
		RequestBody:     nil,
	}
	return data.ToHttpRequest()
}

func BuildTwitchUserRequest(config config.Config, username string) *http.Request {
	data := datamodel.RequestData{
		RequestType:     datamodel.GET,
		RequestURL:      twitchUsersAPI,
		QueryParameters: map[string]string{"login": username},
		Headers:         twitchAuthorizationHeaders(config),
		RequestBody:     nil,
	}
	return data.ToHttpRequest()
}

func BuildTwitchClipsRequest(config config.Config, broadcasterId string) *http.Request {
	data := datamodel.RequestData{
		RequestType:     datamodel.GET,
		RequestURL:      twitchClipsAPI,
		QueryParameters: map[string]string{"broadcaster_id": broadcasterId},
		Headers:         twitchAuthorizationHeaders(config),
		RequestBody:     nil,
	}
	return data.ToHttpRequest()
}

func twitchAuthorizationHeaders(config config.Config) map[string][]string {
	headers := make(map[string][]string)
	headers["Authorization"] = []string{"Bearer " + config.TwitchOAuthConfig.AccessToken}
	headers["Client-Id"] = []string{config.TwitchClientId}
	return headers
}
