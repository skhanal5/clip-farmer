package downloader

import (
	"encoding/json"
	"fmt"
	"github.com/skhanal5/clip-farmer/internal/twitch"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	connectTimeout = 30 * time.Second
	chunkSize      = 1024 * 1024
	downloadDelay  = 5 * time.Second
)

func DownloadClips(path string, clips []twitch.Clip) {
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

// credit to twitch-dl for reference
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
func buildClipDownloadURL(clip twitch.Clip) string {
	log.Print("Making download url for clip with id: " + clip.ID)
	token := clip.PlaybackAccessToken
	var valueTok twitch.Value
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
