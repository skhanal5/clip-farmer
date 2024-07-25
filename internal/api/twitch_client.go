package api

import (
	"errors"
	"io"
	"log"
	"net/http"
)

type TwitchClient struct {
	httpClient *http.Client
}

func NewTwitchClient() *TwitchClient {
	log.Print("Building TwitchClient instance")
	client := &TwitchClient{
		httpClient: &http.Client{},
	}
	return client
}

func (client *TwitchClient) SendGetRequest(req *http.Request) ([]byte, error) {
	resp, err := client.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}

	return body, nil
}
