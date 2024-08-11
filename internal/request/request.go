// Package request contains utility functions to help build a http.Request for downstream API calls.
package request

import (
	"io"
	"net/http"
)

const (
	POST = "POST"
	PUT  = "PUT"
)

// ToHttpRequest contains request metadata and uses them to construct a http.Request.
func ToHttpRequest(requestType string, requestURL string, queryParameters map[string]string, headers map[string][]string, requestBody io.Reader) *http.Request {
	req, _ := http.NewRequest(requestType, requestURL, requestBody)
	if headers != nil {
		req.Header = headers
	}
	queryParams := req.URL.Query()
	for key, value := range queryParameters {
		queryParams.Add(key, value)
	}
	req.URL.RawQuery = queryParams.Encode() // we encode all query params at once here. Do not encode them beforehand.
	return req
}
