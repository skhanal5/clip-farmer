// Package client represents a common http.Client that is used across the application to send
// API to different services.
//
// It exposes a common SendRequest function to handle sending requests and reading the body
package client

import (
	"errors"
	"io"
	"log"
	"net/http"
)

var client = &http.Client{} //setup any config later

// SendRequest handles sending a http.Request and returns the body of the response as a []byte
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
