package tiktok

import (
	"github.com/skhanal5/clip-farmer/internal/request"
	"net/http"
)

func BuildAuthenticationRequest(clientKey string, codeVerifier string) *http.Request {
	const tiktokLoginEndpoint = "https://www.tiktok.com/v2/auth/authorize/"

	queryParams := buildAuthQueryParams(clientKey, codeVerifier)
	headers := buildOAuthAndLoginHeaders()
	return request.ToHttpRequest(request.POST, tiktokLoginEndpoint, queryParams, headers, nil)
}

func buildAuthQueryParams(clientKey string, codeChallenge string) map[string]string {
	const scope = "user.info.basic,video.publish,video.upload"
	const redirectUri = "http://localhost:8080/callback"

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
