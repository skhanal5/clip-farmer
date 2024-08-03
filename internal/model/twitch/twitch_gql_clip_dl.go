package model

type ClipDownloadResponse struct {
	Data       ClipData   `json:"data"`
	Extensions Extensions `json:"extensions"`
}

type ClipData struct {
	Clip Clip `json:"clip"`
}

type Clip struct {
	ID                  string              `json:"id"`
	PlaybackAccessToken PlaybackAccessToken `json:"playbackAccessToken"`
	VideoQualities      []VideoQualities    `json:"videoQualities"`
	TypeName            string              `json:"__typename"`
}

type PlaybackAccessToken struct {
	Signature string `json:"signature"`
	Value     string `json:"value"`
	TypeName  string `json:"__typename"`
}

type Value struct {
	Authorization map[string]any `json:"authorization"`
	ClipURI       string         `json:"clip_uri"`
	ClipSlug      string         `json:"clip_slug"`
	DeviceId      string         `json:"device_id"`
	Expires       int64          `json:"expires"`
	UserId        string         `json:"user_id"`
	Version       int64          `json:"version"`
}

type VideoQualities struct {
	Framerate float32 `json:"framerate"`
	Quality   string  `json:"quality"`
	SourceURL string  `json:"sourceUrl"`
	TypeName  string  `json:"__typename"`
}

type Extensions struct {
	DurationsMS   int    `json:"durationMilliseconds"`
	OperationName string `json:"operationName"`
	RequestId     string `json:"requestId"`
}
