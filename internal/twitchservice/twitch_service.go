package twitchservice

import (
	"encoding/json"
	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/config"
	"github.com/skhanal5/clip-farmer/internal/twitchmodel"
	"github.com/skhanal5/clip-farmer/internal/twitchrequest"
	"log"
)

func FetchTwitchOAuth(config config.Config) twitchmodel.TwitchOAuthResponse {
	oauthRequest := twitchrequest.BuildTwitchOAuthRequest(config)
	log.Print("Invoking Twitch OAuth datamodel")
	responseBody, err := client.SendRequest(oauthRequest)
	if err != nil {
		panic(err)
	}
	var oauthResponse twitchmodel.TwitchOAuthResponse
	err = json.Unmarshal(responseBody, &oauthResponse)
	if err != nil {
		panic(err)
	}
	return oauthResponse
}

func FetchUsers(config config.Config, username string) twitchmodel.TwitchUserResponse {
	userRequest := twitchrequest.BuildTwitchUserRequest(config, username)
	log.Print("Invoking GET Users")
	body, err := client.SendRequest(userRequest)
	if err != nil {
		panic(err)
	}

	var gqlResponse twitchmodel.TwitchUserResponse
	err = json.Unmarshal(body, &gqlResponse)
	if err != nil {
		panic(err)
	}

	return gqlResponse
}

func FetchUserClips(config config.Config, broadcasterId string) twitchmodel.TwitchClipResponse {
	clipsRequest := twitchrequest.BuildTwitchClipsRequest(config, broadcasterId)
	log.Print("Invoking GET Clips")
	responseBody, err := client.SendRequest(clipsRequest)
	if err != nil {
		panic(err)
	}

	var gqlResponse twitchmodel.TwitchClipResponse
	err = json.Unmarshal(responseBody, &gqlResponse)

	if err != nil {
		panic(err)
	}
	return gqlResponse
}
