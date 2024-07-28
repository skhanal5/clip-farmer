package service

import (
	"encoding/json"
	"github.com/skhanal5/clip-farmer/config"
	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/request"
	model "github.com/skhanal5/clip-farmer/model/tiktok"
	"log"
)

func LoginIntoTargetUser(config config.Config) {
	loginRequest := request.BuildTiktokLoginRequest(config)
	log.Print("Invoking TikTok Login request")
	_, err := client.SendRequest(loginRequest)
	if err != nil {
		panic(err)
	}
}

func FetchTikTokOAuth(config config.Config) model.TikTokOAuthResponse {
	loginRequest := request.BuildTiktokLoginRequest(config)
	log.Print("Invoking TikTok Login request")
	responseBody, err := client.SendRequest(loginRequest)
	if err != nil {
		panic(err)
	}
	var oauthResponse model.TikTokOAuthResponse
	err = json.Unmarshal(responseBody, &oauthResponse)
	if err != nil {
		panic(err)
	}
	return oauthResponse
}
