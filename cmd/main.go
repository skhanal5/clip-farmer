package main

import (
	"github.com/skhanal5/clip-farmer/config"
	"github.com/skhanal5/clip-farmer/internal/downloader"
	"github.com/skhanal5/clip-farmer/internal/service"
)

func main() {
	configuration := config.NewConfig()

	token := service.FetchTwitchOAuth(configuration)
	configuration.SetTwitchBearerToken(token)
	user := service.FetchUsers(configuration, "stableronaldo")
	id := user.GetNthUser(0).Id
	resp := service.FetchUserClips(configuration, id)
	err := downloader.DownloadClips(resp.Data)
	if err != nil {
		panic(err)
	}
}
