package twitch

import (
	"bytes"
	"encoding/json"
	"github.com/skhanal5/clip-farmer/internal/request"
	"net/http"
)

func BuildGQLTwitchUserRequest(username string, oauthToken string, clientId string) *http.Request {
	headers := twitchAuthorizationHeadersGQL(oauthToken, clientId)
	requestBody := buildUserReq(username)
	return request.ToHttpRequest(request.POST, twitchGQLAPI, nil, headers, requestBody)
}

func buildUserReq(username string) *bytes.Buffer {
	req := GQLRequest{
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
	jsonData, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(jsonData)
}
