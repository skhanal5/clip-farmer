package main

import (
	"fmt"
	"github.com/skhanal5/clip-farmer/config"
	"github.com/skhanal5/clip-farmer/internal/oauth"
)

func main() {
	configuration := config.LoadConfig()
	fmt.Println(configuration)
	twitchOAuthClient := oauth.NewTwitchClient(configuration)
	oauthResp, err := twitchOAuthClient.FetchAppOAuth()
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Print(oauthResp.AccessToken)
	configuration.SetTwitchClientSecret(oauthResp.AccessToken)
	//twitchService := service.NewTwitchService(configuration)
	//res, err := twitchService.FetchUser("test")
	//fmt.Print(res, err)
}
