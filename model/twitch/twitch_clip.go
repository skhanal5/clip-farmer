package model

type TwitchClip struct {
	Url       string `json:"url"`
	CreatorId string `json:"creator_id"`
	ViewCount int    `json:"view_count"`
	Title     string `json:"title"`
}
