package service

import (
	"encoding/json"
	"github.com/skhanal5/clip-farmer/config"
	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/request"
	model "github.com/skhanal5/clip-farmer/model/twitch"
	"log"
)

func FetchTwitchOAuth(config config.Config) model.TwitchOAuthResponse {
	oauthRequest := request.BuildTwitchOAuthRequest(config)
	log.Print("Invoking Twitch OAuth request")
	responseBody, err := client.SendRequest(oauthRequest)
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

func FetchUsers(config config.Config, username string) model.TwitchUserResponse {
	userRequest := request.BuildTwitchUserRequest(config, username)
	log.Print("Invoking GET Users")
	body, err := client.SendRequest(userRequest)
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

func FetchUserClips(config config.Config, broadcasterId string) model.TwitchClipResponse {
	clipsRequest := request.BuildTwitchClipsRequest(config, broadcasterId)
	log.Print("Invoking GET Clips")
	responseBody, err := client.SendRequest(clipsRequest)
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
