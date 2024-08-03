package request

import (
	"encoding/json"
	"github.com/skhanal5/clip-farmer/internal/config"
	model "github.com/skhanal5/clip-farmer/internal/model/tiktok"
	"github.com/skhanal5/clip-farmer/internal/server"
	"golang.org/x/oauth2"
	"net/http"
)

const (
	tiktokPostEndpoint  = "https://open.tiktokapis.com/v2/post/publish/video/init/"
	tiktokOAuthEndpoint = "https://open.tiktokapis.com/v2/oauth/token/"
	tiktokLoginEndpoint = "https://www.tiktok.com/v2/auth/authorize/"
	redirectUri         = "http://localhost:8080/callback"
	scope               = "user.info.basic,video.publish,video.upload"
	GET                 = "GET"
	POST                = "POST"
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
	queryParams["client_key"] = config.TiktokClientKey
	queryParams["scope"] = scope
	queryParams["response_type"] = server.Code
	queryParams["redirect_uri"] = redirectUri
	queryParams["state"] = "state"
	queryParams["code_challenge"] = codeChallenge
	queryParams["code_challenge_method"] = "S256"
	return queryParams
}
