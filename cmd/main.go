package main

import (
	"github.com/skhanal5/clip-farmer/internal/config"
	"github.com/skhanal5/clip-farmer/internal/downloader"
	"github.com/skhanal5/clip-farmer/internal/service"
)

func main() {
	configuration := config.NewConfig()

	//Delegate the below portion in its own function: TwitchManager/Controller

	token := service.FetchTwitchOAuth(configuration)
	configuration.SetTwitchBearerToken(token)

	user := service.FetchUsers(configuration, "stableronaldo")
	id := user.GetNthUser(0).Id

	resp := service.FetchUserClips(configuration, id)
	err := downloader.DownloadClips(resp.Data)

	if err != nil {
		panic(err)
	}

	//server.LoginIntoTargetUser(configuration)
	//twitchservice.FetchTikTokOAuth(configuration)

}
