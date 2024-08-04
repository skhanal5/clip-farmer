package request

import (
	"io"
	"net/http"
)

func ToHttpRequest(requestType string, requestURL string, queryParameters map[string]string, headers map[string][]string, requestBody io.Reader) *http.Request {
	req, _ := http.NewRequest(requestType, requestURL, requestBody)
	if headers != nil {
		req.Header = headers
	}
	queryParams := req.URL.Query()
	for key, value := range queryParameters {
		queryParams.Add(key, value)
	}
	req.URL.RawQuery = queryParams.Encode()
	return req
}
