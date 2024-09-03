package twitch

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"time"

	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/download"
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
	time.Sleep(requestDelay)
	edges := userRes.Data.User.Clips.Edges
	clips := make([]Clip, 0)
	for _, edge := range edges {
		slug := edge.Node.Slug
		time.Sleep(requestDelay)
		clip := t.fetchClipDownloadInfo(slug)
		clips = append(clips, clip.Data.Clip)
	}
	log.Printf("Fetched %d number of clips", len(clips))
	downloadTwitchClips(user, clips)
}

// fetchUser fetches clip data from the target user and returns it as a UserResponse
func (t *TwitchManager) fetchUser(targetUser string, period string, sort string) UserResponse {
	userRequest := BuildGetClipRequest(targetUser, period, sort, t.clientId, t.clientOAuth)
	log.Println("Fetching user: " + targetUser + " through Twitch GQL API with period: " + period + " sort: " + sort)
	body, err := client.SendRequest(userRequest)
	if err != nil {
		log.Fatal(err)
	}

	var gqlResponse UserResponse
	err = json.Unmarshal(body, &gqlResponse)
	if err != nil { 
		log.Fatal(err)
	}
	return gqlResponse
}


// fetchClipDownloadInfo fetches metadata from the given clip with clipId and returns it as a ClipDownloadResponse
func (t *TwitchManager) fetchClipDownloadInfo(clipId string) ClipDownloadResponse {
	clipsRequest := BuildTwitchClipDownloadRequest(clipId, t.clientId, t.clientOAuth)
	log.Println("Getting clip download info for clip with id: " + clipId + " through Twitch GQL API")
	responseBody, err := client.SendRequest(clipsRequest)
	if err != nil {
		log.Fatal(err)
	}
	var gqlResponse ClipDownloadResponse
	err = json.Unmarshal(responseBody, &gqlResponse)

	if err != nil {
		log.Fatal(err)
	}
	return gqlResponse
}


// downloadTwitchClips allows you to download the specified array of clips onto the
// path in your local filesystem.
func downloadTwitchClips(path string, clips []Clip) {
	const downloadDelay = 20 * time.Second
	makeDownloadOutputDirectory(path)
	for _, clip := range clips {
		mp4Link := constructRawMP4URLFromClip(clip)
		clipOutputPath := path + "/" + clip.ID + ".mp4"
		download.DownloadMP4File(mp4Link, clipOutputPath)
		time.Sleep(downloadDelay)
	}
}

func makeDownloadOutputDirectory(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

// constructRawMP4URLFromClip takes a clip and generates a string containing the mp4 link
// that can be downloaded via a request.
func constructRawMP4URLFromClip(clip Clip) string {
	log.Println("Making download url for clip with id: " + clip.ID)

	token := clip.PlaybackAccessToken
	value := getValueFromToken(clip.PlaybackAccessToken)

	params := url.Values{}
	params.Set("response-content-disposition", "attachment")
	params.Set("sig", token.Signature)
	params.Set("token", token.Value)
	finalURL := fmt.Sprintf("%s?%s", value.ClipURI, params.Encode())
	return finalURL
}

//TODO: This function should not be needed...
// getValueFromToken unmarshals the value field into a Value struct
func getValueFromToken(token PlaybackAccessToken) Value {
	var value Value
	err := json.Unmarshal([]byte(token.Value), &value)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	return value
}