package request

import (
	"github.com/skhanal5/clip-farmer/internal/config"
	"net/http"
	"net/url"
	"strings"
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

func BuildTikTokAuthorizationRequest(config config.Config, codeVerifier string) *http.Request {
	queryParams := loginQueryParameters(config, codeVerifier)
	headers := buildTikTokHeaders()
	return ToHttpRequest(POST, tiktokLoginEndpoint, queryParams, headers, nil)
}

func BuildTikTokOAuthRequest(config config.Config, code string, codeVerifier string) *http.Request {
	oauthBody := buildTikTokOAuthBody(config, code, codeVerifier)
	headers := buildTikTokHeaders()
	return ToHttpRequest(POST, tiktokOAuthEndpoint, make(map[string]string), headers, oauthBody)
}

func buildTikTokOAuthBody(config config.Config, code string, codeVerifier string) *strings.Reader {
	// must be urlencoded
	body := url.Values{}
	body.Add("client_key", config.TiktokClientKey)
	body.Add("client_secret", config.TikTokClientSecret)
	body.Add("code", code)
	body.Add("grant_type", "authorization_code")
	body.Add("redirect_uri", redirectUri)
	body.Add("code_verifier", codeVerifier)
	encoded := strings.NewReader(body.Encode())
	return encoded
}

func buildTikTokHeaders() map[string][]string {
	headers := map[string][]string{}
	headers["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	headers["Cache-Control"] = []string{"no-cache"}
	return headers
}

func loginQueryParameters(config config.Config, codeChallenge string) map[string]string {
	queryParams := make(map[string]string)
	queryParams["client_key"] = config.TiktokClientKey
	queryParams["scope"] = scope
	queryParams["response_type"] = "code"
	queryParams["redirect_uri"] = redirectUri
	queryParams["state"] = "state"
	queryParams["code_challenge"] = codeChallenge
	queryParams["code_challenge_method"] = "S256"
	return queryParams
}
