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
	file, _ := os.Open("clips/stableronaldo/328829385.mp4")
	stat, _ := file.Stat()
	size := stat.Size()
	res2 := manager.UploadVideoAsDraft(configuration, size, file)
	fmt.Println(res2)
}
