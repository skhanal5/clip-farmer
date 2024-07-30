package tiktok

type TikTokOAuthRequestBody struct {
	ClientKey    string `json:"client_key"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	GrantType    string `json:"grant_type"`
	RedirectUri  string `json:"redirect_uri"`
	CodeVerifier string `json:"code_verifier"`
}
