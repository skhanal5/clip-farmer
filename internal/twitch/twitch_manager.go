package twitch

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/skhanal5/clip-farmer/internal/client"
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
	log.Print("Fetching user: " + targetUser + " through Twitch GQL API with period: " + period + " sort: " + sort)
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
	log.Print("Getting clip download info for clip with id: " + clipId + " through Twitch GQL API")
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


const (
	connectTimeout = 30 * time.Second
	chunkSize      = 1024 * 1024
	downloadDelay  = 20 * time.Second
)

// DownloadTwitchClips allows you to download the specified array of clips onto the
// path in your local filesystem.
func downloadTwitchClips(path string, clips []Clip) {
	directoryPath := "clips/" + path
	err := os.MkdirAll(directoryPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	for _, clip := range clips {
		mp4Link := buildClipDownloadURL(clip)
		downloadClip(mp4Link, path+clip.ID, directoryPath)
		time.Sleep(downloadDelay)
	}
}

// downloadClip handles the logic to download a clip given an url, the name of the clip, and a path
// to write the contents of the clip to.
//
// This function was implemented using the logic from twitch-dl. All credit goes to the authors of that library.
func downloadClip(downloadURL string, clipName string, directoryPath string) {
	filepath := directoryPath + "/" + clipName + ".mp4"
	client := http.Client{
		Timeout: connectTimeout,
	}
	resp, err := client.Get(downloadURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	file, err := os.Create(filepath)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	log.Print("Downloading clip " + filepath + " to local filesystem")

	size := int64(0)
	buf := make([]byte, chunkSize)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		_, err = file.Write(buf[:n])
		if err != nil {
			panic(err)
		}
		size += int64(n)
	}
}

// buildClipDownloadURL takes a clip and generates a string containing the mp4 link
// that can be downloaded via a request.
func buildClipDownloadURL(clip Clip) string {
	log.Print("Making download url for clip with id: " + clip.ID)
	token := clip.PlaybackAccessToken
	var valueTok Value
	err := json.Unmarshal([]byte(token.Value), &valueTok)
	if err != nil {
		panic("Error unmarshalling JSON: " + err.Error())
	}

	params := url.Values{}
	params.Set("response-content-disposition", "attachment")
	params.Set("sig", token.Signature)
	params.Set("token", token.Value)

	finalURL := fmt.Sprintf("%s?%s", valueTok.ClipURI, params.Encode())
	return finalURL
}
