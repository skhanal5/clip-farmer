package service

import (
	"encoding/json"
	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/config"
	model "github.com/skhanal5/clip-farmer/internal/model/twitch"
	"github.com/skhanal5/clip-farmer/internal/request"
	"log"
)

func FetchUser(config config.Config, username string) model.UserResponse {
	userRequest := request.BuildGQLTwitchUserRequest(username, config)
	log.Print("Invoking GET Users")
	body, err := client.SendRequest(userRequest)
	if err != nil {
		panic(err)
	}

	var gqlResponse model.UserResponse
	err = json.Unmarshal(body, &gqlResponse)
	if err != nil {
		panic(err)
	}
	return gqlResponse
}

func FetchClipDownloadInfo(config config.Config, clipId string) model.ClipDownloadResponse {
	clipsRequest := request.BuildTwitchClipDownloadRequest(clipId, config)
	log.Print("Invoking GET Clips")
	responseBody, err := client.SendRequest(clipsRequest)
	if err != nil {
		panic(err)
	}

	var gqlResponse model.ClipDownloadResponse
	err = json.Unmarshal(responseBody, &gqlResponse)

	if err != nil {
		panic(err)
	}
	return gqlResponse
}
