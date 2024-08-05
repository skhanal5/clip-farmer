package model

import (
	"encoding/json"
	"log"
	"os"
)

type TikTokOAuthResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	OpenId           string `json:"open_id"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
}

func (o TikTokOAuthResponse) WriteToFile() {
	file, err := os.Create("tiktok_oauth_resp.json")
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	data, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		log.Print(err)
	}
	_, err = file.Write(data)
	if err != nil {
		log.Print(err)
	}
}

func LoadTikTokOAuth() TikTokOAuthResponse {
	var oauthResponse TikTokOAuthResponse
	file, err := os.Open("tiktok_oauth_resp.json")
	if err != nil {
		log.Print(err)
		return oauthResponse
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&oauthResponse)
	if err != nil {
		log.Print(err)
		return oauthResponse
	}

	// Check if the token is expired
	// Get Refresh token
	return oauthResponse
}
