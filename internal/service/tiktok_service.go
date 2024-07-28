package service

import (
	"encoding/json"
	"fmt"
	"github.com/skhanal5/clip-farmer/config"
	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/request"
	"github.com/skhanal5/clip-farmer/internal/server"
	"github.com/skhanal5/clip-farmer/model/tiktok"
	"log"
	"sync"
)

func LoginIntoTargetUser(config config.Config) {
	loginRequest := request.BuildTiktokLoginRequest(config)
	log.Print("Invoking TikTok Login request")
	fmt.Println(loginRequest.URL)
	_, err := client.SendRequest(loginRequest)
	if err != nil {
		panic(err)
	}
	serverDone := &sync.WaitGroup{}
	serverDone.Add(1)
	server.StartServer(serverDone)
	serverDone.Wait()
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
