package client

import (
	"errors"
	"io"
	"log"
	"net/http"
)

var client = &http.Client{} //setup any config later

func SendRequest(req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		log.Print("Received an invalid response " + resp.Status)
		return nil, errors.New(string(body))
	}
	return body, nil
}
