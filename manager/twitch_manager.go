package manager

import (
	"encoding/json"
	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/downloader"
	"github.com/skhanal5/clip-farmer/internal/twitch"
	"log"
	"time"
)

type TwitchManager struct {
	clientId    string
	clientOAuth string
}

func InitTwitchManager(clientId string, clientOAuth string) TwitchManager {
	return TwitchManager{
		clientId:    clientId,
		clientOAuth: clientOAuth,
	}
}

func (t *TwitchManager) FetchAndDownloadClips(user string) {
	const requestDelay = 5 * time.Second // Delay between download attempts

	userRes := fetchUser(t.clientId, t.clientOAuth, user)

	edges := userRes.Data.User.Clips.Edges
	clips := make([]twitch.Clip, 0)
	for _, edge := range edges {
		slug := edge.Node.Slug
		time.Sleep(requestDelay)
		clip := fetchClipDownloadInfo(t.clientId, t.clientOAuth, slug)
		clips = append(clips, clip.Data.Clip)
	}
	downloader.DownloadClips(user, clips)
}

func fetchUser(clientId string, clientOAuth string, targetUser string) twitch.UserResponse {
	userRequest := twitch.BuildGQLTwitchUserRequest(targetUser, clientId, clientOAuth)
	log.Print("Getting user: " + targetUser + " through Twitch GQL API")
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

func fetchClipDownloadInfo(clientId string, clientOAuth string, clipId string) twitch.ClipDownloadResponse {
	clipsRequest := twitch.BuildTwitchClipDownloadRequest(clipId, clientId, clientOAuth)
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
