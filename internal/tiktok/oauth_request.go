package tiktok

import (
	"github.com/skhanal5/clip-farmer/internal/request"
	"net/http"
	"net/url"
	"strings"
)

func BuildOAuthRequest(clientKey string, clientSecret string, code string, codeVerifier string) *http.Request {
	const tiktokOAuthEndpoint = "https://open.tiktokapis.com/v2/oauth/token/"

	oauthBody := buildOAuthBody(clientKey, clientSecret, code, codeVerifier)
	headers := buildOAuthAndLoginHeaders()
	return request.ToHttpRequest(request.POST, tiktokOAuthEndpoint, make(map[string]string), headers, oauthBody)
}

func buildOAuthBody(clientKey string, clientSecret string, code string, codeVerifier string) *strings.Reader {
	const redirectUri = "http://localhost:8080/callback"

	body := url.Values{}
	body.Add("client_key", clientKey)
	body.Add("client_secret", clientSecret)
	body.Add("code", code)
	body.Add("grant_type", "authorization_code")
	body.Add("redirect_uri", redirectUri)
	body.Add("code_verifier", codeVerifier)
	encoded := strings.NewReader(body.Encode())
	return encoded
}

func buildOAuthAndLoginHeaders() map[string][]string {
	headers := map[string][]string{}
	headers["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	headers["Cache-Control"] = []string{"no-cache"}
	return headers
}
