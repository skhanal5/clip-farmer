package main

import (
	"fmt"
	"github.com/skhanal5/clip-farmer/config"
	"github.com/skhanal5/clip-farmer/internal/service"
)

func main() {
	configuration := config.LoadConfig()
	twitchService := service.NewTwitchService(configuration)
	res, err := twitchService.FetchUser("test")
	fmt.Print(res, err)
}
