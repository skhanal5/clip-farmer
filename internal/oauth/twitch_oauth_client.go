package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/skhanal5/clip-farmer/config"
	model "github.com/skhanal5/clip-farmer/model/twitch"
	"io"
	"log"
	"net/http"
)

const (
	twitchOAuthEndpoint = "https://id.twitch.tv/oauth2/token"
)

type TwitchOAuthClient struct {
	headers    map[string]string
	httpClient *http.Client
}

func NewTwitchClient(config config.Config) *TwitchOAuthClient {
	log.Print("Building TwitchOAuthClient instance")
	client := &TwitchOAuthClient{
		headers:    make(map[string]string),
		httpClient: &http.Client{},
	}
	client.headers["client_id"] = config.TwitchClientId
	client.headers["client_secret"] = config.TwitchClientSecret
	return client
}

func (client *TwitchOAuthClient) FetchAppOAuth() (model.TwitchOAuthResponse, error) {
	request, err := http.NewRequest("POST", twitchOAuthEndpoint, nil)
	if err != nil {
		return model.TwitchOAuthResponse{}, err
	}
	queryParams := request.URL.Query()
	fmt.Println(client.headers)
	queryParams.Add("client_id", client.headers["client_id"])
	queryParams.Add("client_secret", client.headers["client_secret"])
	queryParams.Add("grant_type", "client_credentials")
	request.URL.RawQuery = queryParams.Encode()

	resp, err := client.httpClient.Do(request)
	if err != nil {
		return model.TwitchOAuthResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.TwitchOAuthResponse{}, err
	}

	if resp.StatusCode != 200 {
		return model.TwitchOAuthResponse{}, errors.New(string(body))
	}

	var oauthResponse model.TwitchOAuthResponse
	err = json.Unmarshal(body, &oauthResponse)
	if err != nil {
		return model.TwitchOAuthResponse{}, err
	}
	return oauthResponse, nil
}
