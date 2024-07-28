package request

import (
	"encoding/json"
	"github.com/skhanal5/clip-farmer/config"
	model "github.com/skhanal5/clip-farmer/model/tiktok"
	"golang.org/x/oauth2"
	"net/http"
)

const (
	redirectUri = ""
)

const (
	tiktokPostEndpoint  = "https://open.tiktokapis.com/v2/post/publish/video/init/"
	tiktokOAuthEndpoint = "https://open.tiktokapis.com/v2/oauth/token/"
	tiktokLoginEndpoint = "https://www.tiktok.com/v2/auth/authorize/"
)

func BuildTiktokLoginRequest(config config.Config) *http.Request {
	data := RequestData{
		RequestType:     POST,
		RequestURL:      tiktokLoginEndpoint,
		QueryParameters: loginQueryParameters(config),
		Headers:         tiktokLoginHeaders(),
		RequestBody:     nil,
	}
	return data.ToHttpRequest()
}

func BuildTikTokOAuthRequest(config config.Config, code string) *http.Request {
	// must be urlencoded
	body := model.TikTokOAuthRequestBody{
		ClientKey:    config.TiktokClientKey,
		ClientSecret: config.TikTokClientSecret,
		Code:         code,
		GrantType:    "authorization_code",
		RedirectUri:  redirectUri,
		CodeVerifier: oauth2.GenerateVerifier(),
	}

	bodyJson, _ := json.Marshal(body)

	data := RequestData{
		RequestType:     POST,
		RequestURL:      tiktokOAuthEndpoint,
		QueryParameters: nil,
		Headers:         tiktokLoginHeaders(),
		RequestBody:     bodyJson,
	}
	return data.ToHttpRequest()
}

func BuildTikTokContentRequest(config config.Config) {

}

func tiktokLoginHeaders() map[string][]string {
	headers := map[string][]string{}
	headers["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	headers["Cache-Control"] = []string{"no-cache"}
	return headers
}

func loginQueryParameters(config config.Config) map[string]string {
	codeVerifier := oauth2.GenerateVerifier()
	codeChallenge := oauth2.S256ChallengeFromVerifier(codeVerifier)

	queryParams := make(map[string]string)
	queryParams["client_key"] = config.TwitchClientSecret
	queryParams["response_type"] = "code"
	queryParams["redirect_uri"] = "redirect_uri"
	queryParams["state"] = "state"
	queryParams["code_challenge"] = codeChallenge
	queryParams["code_challenge_method"] = "S256"
	return queryParams
}
