package downloader

import (
	"fmt"
	"github.com/skhanal5/clip-farmer/internal/twitchmodel"
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

func DownloadClips(clips []twitchmodel.TwitchClip) error {
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

func convertUrlToMp4(clip string) string {
	index := s.Index(clip, "-preview")
	rawUrl := clip[:index]
	rawUrl += ".mp4"
	return rawUrl
}
