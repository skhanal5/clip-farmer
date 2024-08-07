package tiktok

import (
	"github.com/skhanal5/clip-farmer/internal/request"
	"net/http"
	"os"
)

func BuildVideoUploadRequest(file *os.File, byteRange string, uploadUrl string) *http.Request {
	headers := buildVideoUploadRequestHeaders(byteRange)
	return request.ToHttpRequest(request.PUT, uploadUrl, nil, headers, file)
}

func buildVideoUploadRequestHeaders(byteRange string) map[string][]string {
	headers := map[string][]string{}
	headers["Content-Type"] = []string{"video/mp4"}
	headers["Content-Range"] = []string{byteRange}
	return headers
}
