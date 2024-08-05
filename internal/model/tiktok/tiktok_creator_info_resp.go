package model

type CreatorInfoResponse struct {
	Data  CreatorData `json:"data"`
	Error Error       `json:"error"`
}

type CreatorData struct {
	CreatorAvatarUrl        string   `json:"creator_avatar_url"`
	CreatorUsername         string   `json:"creator_username"`
	PrivacyLevelOptions     []string `json:"privacy_level_options"`
	CommentDisabled         bool     `json:"comment_disabled"`
	DuetDisabled            bool     `json:"duet_disabled"`
	StitchDisabled          bool     `json:"stitch_disabled"`
	MaxVideoPostDurationSec int      `json:"max_video_post_duration_sec"`
}
