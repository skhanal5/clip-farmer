package model

import "errors"

type TwitchGraphQLResponse struct {
	Data   []map[string]interface{} `json:"data"`
	Errors []map[string]interface{} `json:"errors"`
}

func (t *TwitchGraphQLResponse) GetValueFromData(key string) (interface{}, error) {
	if len(t.Data) > 0 {
		firstMap := t.Data[0]
		value, exists := firstMap[key]
		if exists == false {
			return value, errors.New("Key does not exist:" + key)
		}
		return value, nil
	}
	return nil, nil
}

func (t *TwitchGraphQLResponse) HasError() bool {
	if (t.Errors == nil) || (len(t.Errors) == 0) {
		return false
	}
	return true
}
