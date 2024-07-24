package main

import (
	"fmt"
	"github.com/skhanal5/clip-farmer/config"
	"github.com/skhanal5/clip-farmer/internal/oauth"
	"github.com/skhanal5/clip-farmer/internal/service"
)

func main() {
	configuration := config.LoadConfig()
	twitchClient := oauth.NewTwitchOAuthClient(configuration)
	oauthResp, err := twitchClient.FetchAppOAuth()
	if err != nil {
		fmt.Print(err.Error())
	}
	configuration.SetTwitchBearerToken(oauthResp)
	s := service.NewTwitchService(configuration)
	res, err := s.FetchUsers("plaqueboymax")
	id, err := res.GetValueFromData("id")
	idS := fmt.Sprintf("%s", id)
	res, err = s.FetchUserClips(idS)
	fmt.Print(res, err)
}
