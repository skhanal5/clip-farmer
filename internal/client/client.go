package client

import (
	"errors"
	"fmt"
	"io"
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

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}

	fmt.Println(string(body))
	return body, nil
}
