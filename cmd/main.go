package main

import (
	"fmt"
	"github.com/skhanal5/clip-farmer/internal/config"
	"github.com/skhanal5/clip-farmer/manager"
	"os"
)

func main() {
	configuration := config.NewConfig()
	tiktokManager := manager.InitTikTokManager(configuration)

	file, _ := os.Open("clips/stableronaldo/1330804442.mp4")
	stat, _ := file.Stat()
	size := stat.Size()
	res2 := tiktokManager.UploadVideoAsDraft(size, file)
	fmt.Println(res2)
}
