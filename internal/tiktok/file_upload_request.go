package tiktok

import (
	"bytes"
	"encoding/json"
	"github.com/skhanal5/clip-farmer/internal/request"
	"net/http"
)

func BuildFileUploadRequest(oauth string, videoSize int64) *http.Request {
	const fileUploadEndpoint = "https://open.tiktokapis.com/v2/post/publish/inbox/video/init/"

	body := buildFileUploadRequestBody(videoSize, videoSize, 1)
	headers := buildAuthorizationHeaders(oauth)
	return request.ToHttpRequest(request.POST, fileUploadEndpoint, nil, headers, body)
}

func buildFileUploadRequestBody(videoSize int64, chunkSize int64, totalChunkCount int) *bytes.Buffer {
	contents := map[string]interface{}{
		"source":            "FILE_UPLOAD",
		"video_size":        videoSize,
		"chunk_size":        chunkSize,
		"total_chunk_count": totalChunkCount,
	}
	responseBody := map[string]interface{}{
		"source_info": contents,
	}
	jsonBody, err := json.Marshal(responseBody)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(jsonBody)
	return buffer
}

func BuildQueryCreatorInfoRequest(oauth string) *http.Request {
	const tiktokQueryCreatorInfoEndpoint = "https://open.tiktokapis.com/v2/post/publish/creator_info/query/"

	headers := buildAuthorizationHeaders(oauth)
	return request.ToHttpRequest(request.POST, tiktokQueryCreatorInfoEndpoint, nil, headers, nil)
}

func buildAuthorizationHeaders(oauth string) map[string][]string {
	headers := map[string][]string{}
	headers["Authorization"] = []string{"Bearer " + oauth}
	headers["Content-Type"] = []string{"application/json; charset=UTF-8"}
	return headers
}
