package main

import (
	"fmt"
	"github.com/skhanal5/clip-farmer/internal/config"
	"github.com/skhanal5/clip-farmer/manager"
)

func main() {
	configuration := config.NewConfig()
	manager.FetchTiktokOAuth(configuration)
	fmt.Println(configuration)
	res2 := manager.FetchQueryCreatorInfo(configuration)
	fmt.Println(res2)
}
