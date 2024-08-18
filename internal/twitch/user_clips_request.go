package twitch

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/skhanal5/clip-farmer/internal/request"
)

// Convert PERIOD and SORT to Enums, use Optionals
func BuildGetClipRequest(username string, period string, sort string, clientId string, oauthToken string) *http.Request {
	headers := twitchAuthorizationHeaders(clientId, oauthToken)
	requestBody := buildClipReq(username, period, sort)
	return request.ToHttpRequest(request.POST, twitchGQLAPI, nil, headers, requestBody)
}

// convert period and sort to enums
func buildClipReq(username string, period string, sort string) *bytes.Buffer {
	req := buildGQLClipQuery(username, period, sort)
	jsonData, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err )
	}
	return bytes.NewBuffer(jsonData)
}

// convert period and sort to enums
func buildGQLClipQuery(username string, period string, sort string) GQLRequest {
	var req GQLRequest
	if (period != "") {
		req = GQLRequest{
			Query: `query($username: String! $period: ClipsPeriod! $sort: ClipsSort!) {
				user(login: $username) {
					id
					clips (criteria: {period: $period sort: $sort} ) {
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
				"period": period,
				"sort": sort,
			},
		}
	} else {
		req = GQLRequest{
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
	}
	return req
}
