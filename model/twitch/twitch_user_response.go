package model

// TwitchUserResponse is a GraphQL response
// to the Twitch Get Users API
type TwitchUserResponse struct {
	Data   []TwitchUser             `json:"data"`
	Errors []map[string]interface{} `json:"errors"`
}

// TwitchUser represents an individual user in the
// Twitch Get Users API
type TwitchUser struct {
	Id              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	ProfileImageUrl string `json:"profile_image_url"`
	OfflineImageUrl string `json:"offline_image_url"`
	ViewCount       int    `json:"view_count"`
	CreatedAt       string `json:"created_at"`
}

func (t *TwitchUserResponse) GetNthUser(index int) TwitchUser {
	return t.Data[index]
}

func (t *TwitchUserResponse) HasError() bool {
	if (t.Errors == nil) || (len(t.Errors) == 0) {
		return false
	}
	return true
}
