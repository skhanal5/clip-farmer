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

func FetchUserClips(config config.Config, broadcasterId string) twitch.TwitchClipResponse {
	clipsRequest := request.BuildTwitchClipsRequest(config, broadcasterId)
	log.Print("Invoking GET Clips")
	responseBody, err := client.SendRequest(clipsRequest)
	if err != nil {
		panic(err)
	}

	var gqlResponse twitch.TwitchClipResponse
	err = json.Unmarshal(responseBody, &gqlResponse)

	if err != nil {
		panic(err)
	}
	return gqlResponse
}
