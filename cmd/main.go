package main

import (
	"fmt"
	"github.com/skhanal5/clip-farmer/config"
	"github.com/skhanal5/clip-farmer/internal/downloader"
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

	twitchUser, err := s.FetchUsers("plaqueboymax")
	user := twitchUser.GetNthUser(0)

	clips, err := s.FetchUserClips(user.Id)
	downloader.DownloadClips(clips.Data)
}
