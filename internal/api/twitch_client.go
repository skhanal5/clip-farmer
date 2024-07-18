package api

import (
	"github.com/skhanal5/clip-farmer/config"
	"io"
	"log"
	"net/http"
)

type TwitchClient struct {
	headers    map[string]string
	httpClient *http.Client
}

func NewTwitchClient(config config.Config) *TwitchClient {
	log.Print("Building TwitchClient instance")
	client := &TwitchClient{
		headers:    make(map[string]string),
		httpClient: &http.Client{},
	}
	client.headers["Authorization"] = "Bearer " + config.TwitchBearerToken
	client.headers["Client-Id"] = config.TwitchClientId
	return client
}

func (client *TwitchClient) SendGetRequest(endpoint string) ([]byte, error) {
	log.Print("Sending GET request to:" + endpoint)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	setRequestHeaders(req, client.headers)
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func setRequestHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Add(key, value)
	}
}
