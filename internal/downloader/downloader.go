package downloader

import (
	"encoding/json"
	"fmt"
	"github.com/skhanal5/clip-farmer/internal/model/twitch"
	"io"
	"log"
	"net/http"
	"os"
	s "strings"
)

/*
Given a thumbnail url, replace the preview portion
of it with a ".mp4" extension

i.e., https://clips-media-assets2.twitch.tv/jTk1-Xmig5ji1Dll05ivnA/AT-cm%7CjTk1-Xmig5ji1Dll05ivnA-preview-260x147.jpg
*/

func DownloadClips(clips []twitch.TwitchClip) error {
	err := os.Mkdir("clips", os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(clips))
	for _, clip := range clips {
		mp4Link := convertUrlToMp4(clip.ThumbnailURL)
		downloadClip(mp4Link, clip.Id)
	}
	return nil
}

// TODO: Download in a separate goroutine, download only videos >= 30 seconds

func downloadClip(mp4Link string, clipName string) {
	req, _ := http.NewRequest("GET", mp4Link, nil)
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.Header.Get("Content-Type") == "binary/octet-stream" {

		file, err := os.Create("clips/" + clipName + ".mp4")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		log.Print("Downloading clip to local filesystem")

		_, err = io.Copy(file, res.Body)
		if err != nil {
			panic(err)
		}
	}
}

func ParseSourceURL(token twitch.PlaybackAccessToken) {

}

func BuildDownloadLink(token twitch.PlaybackAccessToken) string {
	baseURL := "https://production.assets.clips.twitchcdn.net/AT-cm|%s.mp4?sig=%s\\&token={\"authorization\":{\"forbidden\":false,\"reason\":\"\"},\"clip_uri\":\"\",\"device_id\":\"%s\",\"expires\":%s,\"user_id\":\"%s\",\"version\":2}"

	var valueTok twitch.Value
	value := token.Value
	json.Unmarshal([]byte(value), &valueTok)
	clipId := s.Split(valueTok.ClipURI, "%7")[1]

	fullURL := fmt.Sprintf(baseURL, clipId, token.Signature, valueTok.DeviceId, valueTok.Expires, valueTok.UserId)
	return fullURL
}

func convertUrlToMp4(clip string) string {
	index := s.Index(clip, "-preview")
	rawUrl := clip[:index]
	rawUrl += ".mp4"
	return rawUrl
}
