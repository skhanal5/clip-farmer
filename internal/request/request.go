package request

import (
	"bytes"
	"net/http"
)

const (
	GET  = "GET"
	POST = "POST"
)

// RequestData represents all the request details we will pass in
type RequestData struct {
	RequestType     string
	RequestURL      string
	QueryParameters map[string]string
	Headers         map[string][]string
	RequestBody     []byte
}

//func NewRequestData(requestURL string, requestType string, queryParameters map[string]string, headers map[string][]string) RequestData {
//	return RequestData{
//		requestURL:      requestURL,
//		requestType:     requestType,
//		queryParameters: queryParameters,
//		headers:         headers,
//		requestBody:     make([]byte, 0),
//	}
//}

func (t *RequestData) ToHttpRequest() *http.Request {
	req, _ := http.NewRequest(t.RequestType, t.RequestURL, bytes.NewBuffer(t.RequestBody))
	if t.Headers != nil {
		req.Header = t.Headers
	}
	queryParams := req.URL.Query()
	for key, value := range t.QueryParameters {
		queryParams.Add(key, value)
	}
	req.URL.RawQuery = queryParams.Encode()
	return req
}
