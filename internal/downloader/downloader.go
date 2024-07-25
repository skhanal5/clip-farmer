package downloader

import (
	model "github.com/skhanal5/clip-farmer/model/twitch"
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

func DownloadClips(clips []model.TwitchClip) {
	for _, clip := range clips {
		mp4Link := convertUrlToMp4(clip.ThumbnailURL)
		downloadClip(mp4Link, clip.Id)
	}
}

func downloadClip(mp4Link string, clipName string) {
	req, _ := http.NewRequest("GET", mp4Link, nil)
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	if res.Header.Get("Content-Type") == "video/mp4" {
		// write file to local filesystem
		file, err := os.Create(clipName + ".mp4")
		defer res.Body.Close()
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
