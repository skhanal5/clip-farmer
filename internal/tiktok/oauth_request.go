package tiktok

import (
	"github.com/skhanal5/clip-farmer/internal/request"
	"net/http"
	"net/url"
	"strings"
)

// BuildOAuthRequest defines the TikTok API's oauth request with the given
// secret values, code from an authorization request, and the verifier used to generate the code_challenge in the authorization request.
// Returns a http.Request that populates all necessary fields for this request.
func BuildOAuthRequest(clientKey string, clientSecret string, code string, codeVerifier string) *http.Request {
	const oauthEndpoint = "https://open.tiktokapis.com/v2/oauth/token/"

	oauthBody := buildOAuthBody(clientKey, clientSecret, code, codeVerifier)
	headers := buildOAuthAndLoginHeaders()
	return request.ToHttpRequest(request.POST, oauthEndpoint, make(map[string]string), headers, oauthBody)
}

// buildOAuthBody builds the body of the BuildOAuthRequest and returns a strings.Reader
func buildOAuthBody(clientKey string, clientSecret string, code string, codeVerifier string) *strings.Reader {
	const redirectUri = "http://localhost:8080/callback" // defined in TikTok developer app's callback settings

	body := url.Values{}
	body.Add("client_key", clientKey)
	body.Add("client_secret", clientSecret)
	body.Add("code", code)
	body.Add("grant_type", "authorization_code")
	body.Add("redirect_uri", redirectUri)
	body.Add("code_verifier", codeVerifier)
	encoded := strings.NewReader(body.Encode()) //url encode body values at once
	return encoded
}

// buildOAuthAndLoginHeaders defines common headers used in the TikTok OAuth and Login request phases.
func buildOAuthAndLoginHeaders() map[string][]string {
	headers := map[string][]string{}
	headers["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	headers["Cache-Control"] = []string{"no-cache"}
	return headers
}
