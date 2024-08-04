package main

import (
	"github.com/skhanal5/clip-farmer/internal/config"
	"github.com/skhanal5/clip-farmer/manager"
)

func main() {
	configuration := config.NewConfig()
	manager.FetchAndDownloadClips(configuration)
}
