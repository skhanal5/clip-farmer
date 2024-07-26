package main

import (
	"github.com/skhanal5/clip-farmer/config"
	"github.com/skhanal5/clip-farmer/internal/downloader"
	"github.com/skhanal5/clip-farmer/internal/service"
)

func main() {
	configuration := config.LoadConfig()

	twitchService := service.NewTwitchService()
	token := twitchService.FetchOAuth(configuration)
	configuration.SetTwitchBearerToken(token)
	user := twitchService.FetchUsers("stableronaldo", configuration)
	id := user.GetNthUser(0).Id
	resp := twitchService.FetchUserClips(id, configuration)
	err := downloader.DownloadClips(resp.Data)
	if err != nil {
		panic(err)
	}
}
