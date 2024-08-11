package tiktok

// FileUploadResponse is a struct representing a response to the file upload request
type FileUploadResponse struct {
	Data  FileData `json:"data"`
	Error Error    `json:"error"`
}

// FileData is the data portion of the FileUploadResponse. The values contained
// are used in a video upload request.
type FileData struct {
	PublishId string `json:"publish_id"`
	UploadURL string `json:"upload_url"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	LogId   string `json:"log_id"`
}
