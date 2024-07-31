package main

import (
	"fmt"
	"github.com/skhanal5/clip-farmer/internal/config"
	"github.com/skhanal5/clip-farmer/internal/downloader"
	"github.com/skhanal5/clip-farmer/internal/service"
)

func main() {
	configuration := config.NewConfig()

	//Delegate the below portion in its own function: TwitchManager/Controller

	//token := service.FetchTwitchOAuth(configuration)
	//configuration.SetTwitchBearerToken(token)

	//user := service.FetchUsers(configuration, "stableronaldo")
	//id := user.GetNthUser(0).Id
	//
	//resp := service.FetchUserClips(configuration, id)
	//err := downloader.DownloadClips(resp.Data)
	//
	//if err != nil {
	//	panic(err)
	//}

	user := service.FetchUser(configuration, "stableronaldo")
	slug := user.Data.User.Clips.Edges[0].Node.Slug
	clip := service.FetchClipDownloadInfo(configuration, slug)
	token := clip.Data.Clip.PlaybackAccessToken
	fmt.Println(token)
	fmt.Println(downloader.BuildDownloadLink(token))

	//server.LoginIntoTargetUser(configuration)
	//twitchservice.FetchTikTokOAuth(configuration)

}
