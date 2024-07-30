package server

import (
	"encoding/json"
	"fmt"
	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/config"
	"github.com/skhanal5/clip-farmer/internal/model/request"
	"github.com/skhanal5/clip-farmer/internal/model/tiktok"
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
	StartServer(serverDone)
	serverDone.Wait()
}

func FetchTikTokOAuth(config config.Config) tiktok.TikTokOAuthResponse {
	loginRequest := request.BuildTiktokLoginRequest(config)
	log.Print("Invoking TikTok Login request")
	responseBody, err := client.SendRequest(loginRequest)
	if err != nil {
		panic(err)
	}
	var oauthResponse tiktok.TikTokOAuthResponse
	err = json.Unmarshal(responseBody, &oauthResponse)
	if err != nil {
		panic(err)
	}
	return oauthResponse
}