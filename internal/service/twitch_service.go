package service

import (
	"encoding/json"
	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/config"
	"github.com/skhanal5/clip-farmer/internal/model/request"
	"github.com/skhanal5/clip-farmer/internal/model/twitch"
	"log"
)

func FetchTwitchOAuth(config config.Config) twitch.TwitchOAuthResponse {
	oauthRequest := request.BuildTwitchOAuthRequest(config)
	log.Print("Invoking Twitch OAuth request")
	responseBody, err := client.SendRequest(oauthRequest)
	if err != nil {
		panic(err)
	}
	var oauthResponse twitch.TwitchOAuthResponse
	err = json.Unmarshal(responseBody, &oauthResponse)
	if err != nil {
		panic(err)
	}
	return oauthResponse
}

func FetchUsers(config config.Config, username string) twitch.TwitchUserResponse {
	userRequest := request.BuildTwitchUserRequest(config, username)
	log.Print("Invoking GET Users")
	body, err := client.SendRequest(userRequest)
	if err != nil {
		panic(err)
	}

	var gqlResponse twitch.TwitchUserResponse
	err = json.Unmarshal(body, &gqlResponse)
	if err != nil {
		panic(err)
	}

	return gqlResponse
}

func FetchUser(config config.Config, username string) twitch.UserResponse {
	userRequest := request.BuildGQLTwitchUserRequest(username, config)
	log.Print("Invoking GET Users")
	body, err := client.SendRequest(userRequest)
	if err != nil {
		panic(err)
	}

	var gqlResponse twitch.UserResponse
	err = json.Unmarshal(body, &gqlResponse)
	if err != nil {
		panic(err)
	}
	return gqlResponse
}

func FetchClipDownloadInfo(config config.Config, clipId string) twitch.ClipDownloadResponse {
	clipsRequest := request.BuildTwitchClipDownloadRequest(clipId, config)
	log.Print("Invoking GET Clips")
	responseBody, err := client.SendRequest(clipsRequest)
	if err != nil {
		panic(err)
	}

	var gqlResponse twitch.ClipDownloadResponse
	err = json.Unmarshal(responseBody, &gqlResponse)

	if err != nil {
		panic(err)
	}
	return gqlResponse
}
