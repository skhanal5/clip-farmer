package tiktok

import (
	"encoding/json"
	"fmt"
	"github.com/skhanal5/clip-farmer/internal/client"
	"log"
	"os"
	"time"
)

// TikTokManager contains all the necessary secret values to interact with the TikTok account via
// the TikTok API
type TikTokManager struct {
	oauthToken string
}

// InitTikTokManager Initializes a TikTokManager with an oauth token and returns an instance of it
func InitTikTokManager(oauthToken string) TikTokManager {
	return TikTokManager{oauthToken: oauthToken}
}

// UploadVideos uploads all videos in the directory specified onto the TikTok account
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

// UploadVideo uploads the specified video in the filePath onto the TikTok account
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
	t.uploadVideoAsDraft(stat.Size(), file)
}

// uploadVideoAsDraft uploads the corresponding file as a draft onto the TikTok account. Returns the response body from the
// corresponding API call.
func (t *TikTokManager) uploadVideoAsDraft(size int64, file *os.File) string {
	if size > 64000000 {
		panic("file size too big to be uploaded in one chunk")
	}
	response := t.sendFileUploadReq(size)
	return t.sendVideoUploadReq(file, size, response)
}

// sendFileUploadReq sends a request to allow uploading a video of the specified size to TikTok's API and returns the response from that call
// as a FileUploadResponse struct. This must be invoked before sendVideoUploadReq to initiate an upload request.
func (t *TikTokManager) sendFileUploadReq(size int64) FileUploadResponse {
	fileUploadReq := BuildFileUploadRequest(t.oauthToken, size)
	res, err := client.SendRequest(fileUploadReq)
	if err != nil {
		panic(err)
	}
	var videoUploadRes FileUploadResponse
	err = json.Unmarshal(res, &videoUploadRes)
	if err != nil {
		panic(err)
	}
	return videoUploadRes
}

// sendVideoUploadReq sends a request to upload the video represented by the specified file and size. In addition, it takes
// in a FileUploadResponse from a sendFileUploadReq call which is a pre-requisite when uploading onto a TikTok account via an API.
// Returns a string representing the body of the video upload request.
func (t *TikTokManager) sendVideoUploadReq(file *os.File, size int64, response FileUploadResponse) string {
	byteRange := fmt.Sprintf("bytes 0-%d/%d", size-1, size)
	videoUploadReq := BuildVideoUploadRequest(file, byteRange, response.Data.UploadURL)
	res, err := client.SendRequest(videoUploadReq)
	if err != nil {
		panic(err)
	}
	return string(res)
}
