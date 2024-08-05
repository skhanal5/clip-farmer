package request

import (
	"bytes"
	"encoding/json"
	"github.com/skhanal5/clip-farmer/internal/config"
	model "github.com/skhanal5/clip-farmer/internal/model/twitch"
	"net/http"
)

const (
	twitchGQLAPI = "https://gql.twitch.tv/gql"
)

func BuildTwitchClipDownloadRequest(clipSlug string, config config.Config) *http.Request {
	return toHttpRequest(POST, twitchGQLAPI, nil, twitchAuthorizationHeadersGQL(config), buildGQLClipsRequestBody(clipSlug))
}

func BuildGQLTwitchUserRequest(username string, config config.Config) *http.Request {
	return toHttpRequest(POST, twitchGQLAPI, nil, twitchAuthorizationHeadersGQL(config), buildUserReq(username))
}

func buildUserReq(username string) *bytes.Buffer {
	request := model.GQLRequest{
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
	return bytes.NewBuffer(jsonData)
}

func buildGQLClipsRequestBody(clipSlug string) *bytes.Buffer {
	request := model.PersistedGQLRequest{
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
	return bytes.NewBuffer(jsonData)
}

func twitchAuthorizationHeadersGQL(config config.Config) map[string][]string {
	headers := make(map[string][]string)
	headers["Authorization"] = []string{"OAuth" + config.TwitchClientOAuth}
	headers["Client-Id"] = []string{config.TwitchClientId}
	return headers
}
