package main

import (
	"github.com/skhanal5/clip-farmer/internal/config"
	"github.com/skhanal5/clip-farmer/internal/manager"
)

func main() {
	configuration := config.NewConfig()
	username := "jasontheween"
	manager.FetchAndDownloadClips(configuration, username)
}
