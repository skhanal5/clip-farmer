package request

import (
	"encoding/json"
	"github.com/skhanal5/clip-farmer/internal/config"
	"net/http"
)

const (
	twitchGQLAPI = "https://gql.twitch.tv/gql"
)

type GQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type PersistedGQLRequest struct {
	OperationName string         `json:"operationName"`
	Variables     QueryVariables `json:"variables"`
	Query         string         `json:"query"`
	Extensions    Extensions     `json:"extensions"`
}

type QueryVariables struct {
	Slug string `json:"slug"`
}

type Extensions struct {
	PersistedQuery PersistedQuery `json:"persistedQuery"`
}

type PersistedQuery struct {
	Version    int    `json:"version"`
	Sha256Hash string `json:"sha256Hash"`
}

type VideoAccessTokenClipRequest struct {
	OperationName string         `json:"operationName"`
	Variables     QueryVariables `json:"variables"`
	Query         string         `json:"query"`
	Extensions    Extensions     `json:"extensions"`
}

type ClipRequest struct {
}

func BuildTwitchClipDownloadRequest(clipSlug string, config config.Config) *http.Request {
	data := RequestData{
		RequestType:     POST,
		RequestURL:      twitchGQLAPI,
		QueryParameters: nil,
		Headers:         twitchAuthorizationHeadersGQL(config),
		RequestBody:     buildGQLClipsRequestBody(clipSlug),
	}
	return data.ToHttpRequest()
}

func BuildGQLTwitchUserRequest(username string, config config.Config) *http.Request {
	data := RequestData{
		RequestType:     POST,
		RequestURL:      twitchGQLAPI,
		QueryParameters: nil,
		Headers:         twitchAuthorizationHeadersGQL(config),
		RequestBody:     buildUserReq(username),
	}
	return data.ToHttpRequest()
}

func buildUserReq(username string) []byte {
	request := GQLRequest{
		Query: `query($username: String!) {
			user(login: $username) {
				id
				clips {
					edges {
						node {
							slug	
						}
					}
				}
			}
		}`,
		Variables: map[string]interface{}{
			"username": username,
		},
	}
	jsonData, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	return jsonData
}

func buildGQLClipsRequestBody(clipSlug string) []byte {
	request := PersistedGQLRequest{
		OperationName: "VideoAccessToken_Clip",
		Query:         "query VideoAccessToken_Clip($slug: ID!) { clip(slug: $slug) { broadcaster { displayName } createdAt curator { displayName id } durationSeconds id tiny: thumbnailURL(width: 86, height: 45) small: thumbnailURL(width: 260, height: 147) medium: thumbnailURL(width: 480, height: 272) title videoQualities { frameRate quality sourceURL } viewCount } }",
	}

	request.Variables.Slug = clipSlug
	request.Extensions.PersistedQuery.Version = 1
	request.Extensions.PersistedQuery.Sha256Hash = "36b89d2507fce29e5ca551df756d27c1cfe079e2609642b4390aa4c35796eb11"

	jsonData, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	return jsonData
}

func twitchAuthorizationHeadersGQL(config config.Config) map[string][]string {
	headers := make(map[string][]string)
	headers["Authorization"] = []string{"OAuth" + config.TwitchClientOAuth}
	headers["Client-Id"] = []string{config.TwitchClientId}
	return headers
}
