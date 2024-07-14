package model

type TwitchUser struct {
	ID              string `json:"id"`
	DisplayName     string `json:"display_name"`
	ProfileImageUrl string `json:"profile_image_url"`
	ViewCount       int    `json:"view_count"`
}
