package tiktokservice

import (
	"encoding/json"
	"fmt"
	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/config"
	"github.com/skhanal5/clip-farmer/internal/server"
	"github.com/skhanal5/clip-farmer/internal/tiktokmodel"
	"github.com/skhanal5/clip-farmer/internal/tiktokrequest"
	"log"
	"sync"
)

func LoginIntoTargetUser(config config.Config) {
	loginRequest := tiktokrequest.BuildTiktokLoginRequest(config)
	log.Print("Invoking TikTok Login datamodel")
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

func FetchTikTokOAuth(config config.Config) tiktokmodel.TikTokOAuthResponse {
	loginRequest := tiktokrequest.BuildTiktokLoginRequest(config)
	log.Print("Invoking TikTok Login datamodel")
	responseBody, err := client.SendRequest(loginRequest)
	if err != nil {
		panic(err)
	}
	var oauthResponse tiktokmodel.TikTokOAuthResponse
	err = json.Unmarshal(responseBody, &oauthResponse)
	if err != nil {
		panic(err)
	}
	return oauthResponse
}
