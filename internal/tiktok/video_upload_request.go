package tiktok

import (
	"github.com/skhanal5/clip-farmer/internal/request"
	"net/http"
	"os"
)

// BuildVideoUploadRequest defines the TikTok API's video upload request with the given
// file, byteRange, and uploadURL. byteRange must be consistent with the videoSize passed in the BuildFileUploadRequest.
// Assumes we will send all bytes of the video at once.
func BuildVideoUploadRequest(file *os.File, byteRange string, uploadUrl string) *http.Request {
	headers := buildVideoUploadRequestHeaders(byteRange)
	return request.ToHttpRequest(request.PUT, uploadUrl, nil, headers, file)
}

// buildVideoUploadRequestHeaders defines the headers passed into
// the BuildVideoUploadRequest
func buildVideoUploadRequestHeaders(byteRange string) map[string][]string {
	headers := map[string][]string{}
	headers["Content-Type"] = []string{"video/mp4"}
	headers["Content-Range"] = []string{byteRange}
	return headers
}
