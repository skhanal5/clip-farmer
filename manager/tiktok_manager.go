package manager

import (
	"encoding/json"
	"fmt"
	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/tiktok"
	"log"
	"os"
	"time"
)

type TikTokManager struct {
	oauthToken string
}

func InitTikTokManager(oauthToken string) TikTokManager {
	return TikTokManager{oauthToken: oauthToken}
}

func (t *TikTokManager) UploadVideos(directory string) {
	dir, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dir {
		fileInfo, err := file.Info()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)
		t.UploadVideo(directory + "/" + fileInfo.Name())
	}
}

func (t *TikTokManager) UploadVideo(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		log.Fatalf("Tried to open file properties, received error: %s", err)
	}
	if stat.IsDir() {
		log.Fatalf("Path to file: %s is not a valid file", filepath)
	}
	uploadVideoAsDraft(stat.Size(), file, t.oauthToken)
}

func uploadVideoAsDraft(size int64, file *os.File, accessToken string) string {
	if size > 64000000 {
		panic("file size too big to be uploaded in one chunk")
	}
	response := sendFileUploadReq(accessToken, size)
	return sendVideoUploadReq(file, size, response)
}

func sendFileUploadReq(accessToken string, size int64) tiktok.FileUploadResponse {
	fileUploadReq := tiktok.BuildFileUploadRequest(accessToken, size)
	res, err := client.SendRequest(fileUploadReq)
	if err != nil {
		panic(err)
	}
	var videoUploadRes tiktok.FileUploadResponse
	err = json.Unmarshal(res, &videoUploadRes)
	if err != nil {
		panic(err)
	}
	return videoUploadRes
}

func sendVideoUploadReq(file *os.File, size int64, response tiktok.FileUploadResponse) string {
	byteRange := fmt.Sprintf("bytes 0-%d/%d", size-1, size)
	videoUploadReq := tiktok.BuildVideoUploadRequest(file, byteRange, response.Data.UploadURL)
	res, err := client.SendRequest(videoUploadReq)
	if err != nil {
		panic(err)
	}
	return string(res)
}
