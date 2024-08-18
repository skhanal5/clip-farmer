package manager

import (
	"encoding/json"
	"log"
	"time"

	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/downloader"
	"github.com/skhanal5/clip-farmer/internal/twitch"
)

// TwitchManager contains all necessary secret values required to interact with Twitch's API
type TwitchManager struct {
	clientId    string
	clientOAuth string
}

// InitTwitchManager defines a TwitchManager with the secret values passed in
func InitTwitchManager(clientId string, clientOAuth string) TwitchManager {
	return TwitchManager{
		clientId:    clientId,
		clientOAuth: clientOAuth,
	}
}

// FetchAndDownloadClips retrieves the clips from the user and downloads it into the local filesystem
func (t *TwitchManager) FetchAndDownloadClips(user string, period string, sort string) {
	const requestDelay = 5 * time.Second // Delay between download attempts

	userRes := t.fetchUser(user, period, sort)

	edges := userRes.Data.User.Clips.Edges
	clips := make([]twitch.Clip, 0)
	for _, edge := range edges {
		slug := edge.Node.Slug
		time.Sleep(requestDelay)
		clip := t.fetchClipDownloadInfo(slug)
		clips = append(clips, clip.Data.Clip)
	}
	log.Printf("Fetched %d number of clips", len(clips))
	downloader.DownloadTwitchClips(user, clips)
}

// fetchUser fetches clip data from the target user and returns it as a UserResponse
func (t *TwitchManager) fetchUser(targetUser string, period string, sort string) twitch.UserResponse {
	userRequest := twitch.BuildGetClipRequest(targetUser, period, sort, t.clientId, t.clientOAuth)
	log.Print("Getting user: " + targetUser + " through Twitch GQL API with period: " + period + " sort: " + sort)
	body, err := client.SendRequest(userRequest)
	if err != nil {
		log.Fatal(err)
	}

	var gqlResponse twitch.UserResponse
	err = json.Unmarshal(body, &gqlResponse)
	if err != nil {
		log.Fatal(err)
	}
	return gqlResponse
}


// fetchClipDownloadInfo fetches metadata from the given clip with clipId and returns it as a ClipDownloadResponse
func (t *TwitchManager) fetchClipDownloadInfo(clipId string) twitch.ClipDownloadResponse {
	clipsRequest := twitch.BuildTwitchClipDownloadRequest(clipId, t.clientId, t.clientOAuth)
	log.Print("Getting clip download info for clip with id: " + clipId + " through Twitch GQL API")
	responseBody, err := client.SendRequest(clipsRequest)
	if err != nil {
		log.Fatal(err)
	}
	var gqlResponse twitch.ClipDownloadResponse
	err = json.Unmarshal(responseBody, &gqlResponse)

	if err != nil {
		log.Fatal(err)
	}
	return gqlResponse
}
