package twitch

import (
	"bytes"
	"encoding/json"
	"github.com/skhanal5/clip-farmer/internal/request"
	"net/http"
)

const (
	twitchGQLAPI = "https://gql.twitch.tv/gql"
)

func BuildTwitchClipDownloadRequest(clipSlug string, clientId string, oauthToken string) *http.Request {
	headers := twitchAuthorizationHeadersGQL(clientId, oauthToken)
	requestBody := buildGQLClipsRequestBody(clipSlug)
	return request.ToHttpRequest(request.POST, twitchGQLAPI, nil, headers, requestBody)
}

func buildGQLClipsRequestBody(clipSlug string) *bytes.Buffer {
	req := PersistedGQLRequest{
		OperationName: "VideoAccessToken_Clip",
		Query:         "query VideoAccessToken_Clip($slug: ID!) { clip(slug: $slug) { broadcaster { displayName } createdAt curator { displayName id } durationSeconds id tiny: thumbnailURL(width: 86, height: 45) small: thumbnailURL(width: 260, height: 147) medium: thumbnailURL(width: 480, height: 272) title videoQualities { frameRate quality sourceURL } viewCount } }",
	}

	req.Variables.Slug = clipSlug
	req.Extensions.PersistedQuery.Version = 1
	req.Extensions.PersistedQuery.Sha256Hash = "36b89d2507fce29e5ca551df756d27c1cfe079e2609642b4390aa4c35796eb11"

	jsonData, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(jsonData)
}

func twitchAuthorizationHeadersGQL(clientId string, oauthToken string) map[string][]string {
	headers := make(map[string][]string)
	headers["Authorization"] = []string{"OAuth" + oauthToken}
	headers["Client-Id"] = []string{clientId}
	return headers
}
