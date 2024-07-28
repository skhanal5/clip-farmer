package twitchmodel

// TwitchClipResponse is a GraphQL response
// to the Twitch Get Clips API
type TwitchClipResponse struct {
	Data   []TwitchClip             `json:"data"`
	Errors []map[string]interface{} `json:"errors"`
}

// TwitchClip represents an individual twitchrequest clip in the
// Twitch Get Clips API response
type TwitchClip struct {
	Id              string  `json:"id"`
	URL             string  `json:"url"`
	EmbedURL        string  `json:"embed_url"`
	BroadcasterID   string  `json:"broadcaster_id"`
	BroadcasterName string  `json:"broadcaster_name"`
	CreatorID       string  `json:"creator_id"`
	CreatorName     string  `json:"creator_name"`
	VideoID         string  `json:"video_id"`
	GameID          string  `json:"game_id"`
	Language        string  `json:"language"`
	Title           string  `json:"title"`
	ViewCount       int     `json:"view_count"`
	CreatedAt       string  `json:"created_at"`
	ThumbnailURL    string  `json:"thumbnail_url"`
	Duration        float64 `json:"duration"`
	VODOffset       int     `json:"vod_offset"`
}

func (t *TwitchClipResponse) GetNthClip(index int) TwitchClip {
	return t.Data[index]
}

func (t *TwitchClipResponse) HasError() bool {
	if (t.Errors == nil) || (len(t.Errors) == 0) {
		return false
	}
	return true
}
