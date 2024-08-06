package main

import (
	"fmt"
	"github.com/skhanal5/clip-farmer/internal/config"
	"github.com/skhanal5/clip-farmer/manager"
	"os"
)

func main() {
	configuration := config.NewConfig()
	manager.FetchTiktokOAuth(configuration)
	fmt.Println(configuration)
	file, _ := os.Stat("clips/stableronaldo/328829385.mp4")
	res2 := manager.UploadVideoAsDraft(configuration, file)
	fmt.Println(res2)
}
