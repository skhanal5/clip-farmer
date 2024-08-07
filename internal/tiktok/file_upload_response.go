package tiktok

type FileUploadResponse struct {
	Data  FileData `json:"data"`
	Error Error    `json:"error"`
}

type FileData struct {
	PublishId string `json:"publish_id"`
	UploadURL string `json:"upload_url"`
}
