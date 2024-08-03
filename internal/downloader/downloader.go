package downloader

import (
	"encoding/json"
	"fmt"
	model "github.com/skhanal5/clip-farmer/internal/model/twitch"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	connectTimeout = 30 * time.Second
	chunkSize      = 1024 * 1024 // 1 MB chunk size
	downloadDelay  = 5 * time.Second
)

func DownloadClips(username string, clips []model.Clip) error {
	path := "clips/" + username
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
	for _, clip := range clips {
		mp4Link := buildDownloadLink(clip)
		fmt.Println(mp4Link)
		//downloadClip(mp4Link, clip.ID, path)
		time.Sleep(downloadDelay)
	}
	return nil
}

func downloadClip(downloadURL string, clipName string, path string) {
	filepath := path + clipName + ".mp4"
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

	log.Print("Downloading clip to local filesystem")

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
func buildDownloadLink(clip model.Clip) string {

	token := clip.PlaybackAccessToken
	var valueTok model.Value
	err := json.Unmarshal([]byte(token.Value), &valueTok)
	if err != nil {
		panic("Error unmarshaling JSON: " + err.Error())
	}

	// Build URL parameters
	params := url.Values{}
	params.Set("response-content-disposition", "attachment")
	params.Set("sig", token.Signature)
	params.Set("token", token.Value)

	// Construct the final URL
	finalURL := fmt.Sprintf("%s?%s", valueTok.ClipURI, params.Encode())
	return finalURL
}
