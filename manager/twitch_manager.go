package manager

import (
	"encoding/json"
	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/config"
	"github.com/skhanal5/clip-farmer/internal/downloader"
	"github.com/skhanal5/clip-farmer/internal/twitch"
	"log"
	"time"
)

const (
	requestDelay = 5 * time.Second // Delay between download attempts
)

func FetchAndDownloadClips(config config.Config) {
	user := fetchUser(config)
	edges := user.Data.User.Clips.Edges
	clips := make([]twitch.Clip, 0)
	for _, edge := range edges {
		slug := edge.Node.Slug
		time.Sleep(requestDelay)
		clip := fetchClipDownloadInfo(config, slug)
		clips = append(clips, clip.Data.Clip)
	}
	downloader.DownloadClips(config.TwitchTargetCreator, clips)
}

func fetchUser(config config.Config) twitch.UserResponse {
	userRequest := twitch.BuildGQLTwitchUserRequest(config.TwitchTargetCreator, config.TwitchClientOAuth, config.TwitchClientId)
	log.Print("Getting user: " + config.TwitchTargetCreator + " through Twitch GQL API")
	body, err := client.SendRequest(userRequest)
	if err != nil {
		panic(err)
	}

	var gqlResponse twitch.UserResponse
	err = json.Unmarshal(body, &gqlResponse)
	if err != nil {
		panic(err)
	}
	return gqlResponse
}

func fetchClipDownloadInfo(config config.Config, clipId string) twitch.ClipDownloadResponse {
	clipsRequest := twitch.BuildTwitchClipDownloadRequest(clipId, config.TwitchClientOAuth, config.TwitchClientId)
	log.Print("Getting clip download info for clip with id: " + clipId + " through Twitch GQL API")
	responseBody, err := client.SendRequest(clipsRequest)
	if err != nil {
		panic(err)
	}

	var gqlResponse twitch.ClipDownloadResponse
	err = json.Unmarshal(responseBody, &gqlResponse)

	if err != nil {
		panic(err)
	}
	return gqlResponse
}