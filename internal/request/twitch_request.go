package request

import (
	"net/http"
)

type TwitchRequestData struct {
	RequestURL  string
	RequestType string
	Query       map[string]string
	Headers     map[string][]string
}

type TwitchRequest struct {
	Request *http.Request
}

func NewRequest(requestURL string, requestType string, query map[string]string, headers map[string][]string) TwitchRequestData {
	return TwitchRequestData{
		RequestURL:  requestURL,
		RequestType: requestType,
		Query:       query,
		Headers:     headers,
	}
}

func (t *TwitchRequestData) BuildRequest() TwitchRequest {
	req, _ := http.NewRequest(t.RequestType, t.RequestURL, nil)
	if t.Headers != nil {
		req.Header = t.Headers
	}
	queryParams := req.URL.Query()
	for key, value := range t.Query {
		queryParams.Add(key, value)
	}
	req.URL.RawQuery = queryParams.Encode()
	return TwitchRequest{
		Request: req,
	}
}
