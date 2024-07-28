package main

import (
	"github.com/skhanal5/clip-farmer/internal/config"
	"github.com/skhanal5/clip-farmer/internal/tiktokservice"
)

func main() {
	configuration := config.NewConfig()

	//Delegate the below portion in its own function: TwitchManager/Controller

	//token := twitchservice.FetchTwitchOAuth(configuration)
	//configuration.SetTwitchBearerToken(token)
	//
	//user := twitchservice.FetchUsers(configuration, "stableronaldo")
	//id := user.GetNthUser(0).Id
	//
	//resp := twitchservice.FetchUserClips(configuration, id)
	//err := downloader.DownloadClips(resp.Data)
	//
	//if err != nil {
	//	panic(err)
	//}

	tiktokservice.LoginIntoTargetUser(configuration)
	//twitchservice.FetchTikTokOAuth(configuration)

}
