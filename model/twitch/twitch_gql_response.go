package model

type TwitchGraphQLResponse struct {
	Data   Data    `json:"data"`
	Errors []Error `json:"errors"`
}

type Data struct {
	Data map[string]interface{}
}

type Error struct {
	Error map[string]interface{}
}
