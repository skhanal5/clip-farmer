package tiktok

import (
	"github.com/skhanal5/clip-farmer/internal/request"
	"net/http"
)

// BuildAuthenticationRequest defines the TikTok API's authentication request with the given
// clientKey and codeVerifier. Returns a http.Request with all request values preconfigured.
func BuildAuthenticationRequest(clientKey string, codeVerifier string) *http.Request {
	const loginEndpoint = "https://www.tiktok.com/v2/auth/authorize/"

	queryParams := buildAuthQueryParams(clientKey, codeVerifier)
	headers := buildOAuthAndLoginHeaders()
	return request.ToHttpRequest(request.POST, loginEndpoint, queryParams, headers, nil)
}

// buildAuthQueryParams defines all query params that are needed for the auth request and
// returns them as a map[string]string
func buildAuthQueryParams(clientKey string, codeChallenge string) map[string]string {
	const scope = "user.info.basic,video.publish,video.upload" //needed scope for posting
	const redirectUri = "http://localhost:8080/callback"       // defined in your TikTok app settings

	queryParams := make(map[string]string)
	queryParams["client_key"] = clientKey
	queryParams["scope"] = scope
	queryParams["response_type"] = "code"
	queryParams["redirect_uri"] = redirectUri
	queryParams["state"] = "state"
	queryParams["code_challenge"] = codeChallenge
	queryParams["code_challenge_method"] = "S256"
	return queryParams
}
