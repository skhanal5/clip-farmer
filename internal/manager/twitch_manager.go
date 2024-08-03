package manager

import (
	"github.com/skhanal5/clip-farmer/internal/config"
	"github.com/skhanal5/clip-farmer/internal/downloader"
	model "github.com/skhanal5/clip-farmer/internal/model/twitch"
	"github.com/skhanal5/clip-farmer/internal/service"
	"time"
)

const (
	requestDelay = 5 * time.Second // Delay between download attempts
)

func FetchAndDownloadClips(config config.Config, username string) {
	user := service.FetchUser(config, username)
	edges := user.Data.User.Clips.Edges
	clips := make([]model.Clip, 0)
	for _, edge := range edges {
		slug := edge.Node.Slug
		time.Sleep(requestDelay)
		clip := service.FetchClipDownloadInfo(config, slug)
		clips = append(clips, clip.Data.Clip)
	}
	downloader.DownloadClips(username, clips)
}
