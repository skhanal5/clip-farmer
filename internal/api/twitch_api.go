package api

import (
	"github.com/skhanal5/clip-farmer/config"
	"github.com/skhanal5/clip-farmer/model/twitch"
	"net/http"
)

const twitchClipsAPI = "https://api.twitch.tv/helix/clips"
const twitchUsersAPI = "https://api.twitch.tv/helix/users"

func GetUser(username string, config config.Config) (model.TwitchUser, error) {
	client := &http.Client{}
	apiURL := getTwitchUsersEndpoint(username)

	req, err := buildGetRequest(apiURL, config)
	if err != nil {
		return model.TwitchUser{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return model.TwitchUser{}, err
	}
	defer resp.Body.Close()

}

func GetTwitchClips(broadcasterId string, config config.Config) (model.TwitchClip, error) {
	client := &http.Client{}
	apiURL := getTwitchClipsEndpoint(broadcasterId)

	req, err := buildGetRequest(apiURL, config)
	if err != nil {
		return model.TwitchClip{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return model.TwitchClip{}, err
	}

	defer resp.Body.Close()
	return resp, nil
}

func buildGetRequest(url string, config config.Config) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+config.TwitchBearerToken)
	req.Header.Add("Client-Id", config.TwitchClientId)
	return req, nil
}

func getTwitchClipsEndpoint(broadcasterId string) string {
	return twitchClipsAPI + "?broadcaster_id=" + broadcasterId
}

func getTwitchUsersEndpoint(loginName string) string {
	return twitchUsersAPI + "?login_name=" + loginName
}
