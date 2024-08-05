package manager

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/config"
	model "github.com/skhanal5/clip-farmer/internal/model/tiktok"
	"github.com/skhanal5/clip-farmer/internal/request"
	"github.com/skhanal5/clip-farmer/internal/server"
	"log"
	"math/big"
	"strings"
	"sync"
)

func FetchQueryCreatorInfo(config config.Config) model.CreatorInfoResponse {
	creatorInfoReq := request.BuildTikTokQueryCreatorInfoRequest(config)
	fmt.Println(creatorInfoReq)
	res, err := client.SendRequest(creatorInfoReq)
	if err != nil {
		panic(err)
	}
	var creatorInfoRes model.CreatorInfoResponse
	fmt.Println(string(res))
	err = json.Unmarshal(res, &creatorInfoRes)
	if err != nil {
		panic(err)
	}
	return creatorInfoRes
}

func FetchTiktokOAuth(config config.Config) {
	if config.TikTokOAuth.AccessToken != "" {
		return
	}
	codeVerifier := generateCodeVerifier(64)
	codeChallenge := generateRawCodeChallenge(codeVerifier)
	code := authenticateTikTokUser(config, codeChallenge)
	oauth := sendTikTokOAuthRequest(config, code, codeVerifier)
	config.SetTikTokOAuth(oauth)
	oauth.WriteToFile()
}

func generateCodeVerifier(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~"
	var result strings.Builder
	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		result.WriteByte(chars[index.Int64()])
	}
	return result.String()
}

func generateRawCodeChallenge(codeVerifier string) string {
	hash := sha256.New()
	hash.Write([]byte(codeVerifier))
	hashBytes := hash.Sum(nil)
	return fmt.Sprintf("%x", hashBytes)
}

func authenticateTikTokUser(config config.Config, codeChallenge string) string {
	authRequest := request.BuildTikTokAuthorizationRequest(config, codeChallenge)
	log.Print("Invoking TikTok Login request")
	fmt.Println(authRequest.URL.String())
	codeChan := make(chan string)
	serverDone := &sync.WaitGroup{}
	serverDone.Add(1)
	server.StartCallbackServer(serverDone, codeChan)
	value := <-codeChan
	return value
}

func sendTikTokOAuthRequest(config config.Config, code string, codeVerifier string) model.TikTokOAuthResponse {
	oauthReq := request.BuildTikTokOAuthRequest(config, code, codeVerifier)
	log.Print("Invoking TikTok OAuth request")
	responseBody, err := client.SendRequest(oauthReq)
	if err != nil {
		panic(err)
	}
	var oauthResponse model.TikTokOAuthResponse
	fmt.Println(string(responseBody))
	err = json.Unmarshal(responseBody, &oauthResponse)
	if err != nil {
		panic(err)
	}
	return oauthResponse
}
