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

func FetchTiktokOAuth(config config.Config) model.TikTokOAuthResponse {
	codeVerifier := generateCodeVerifier(64)
	codeChallenge := generateRawCodeChallenge(codeVerifier)
	code := authenticateTikTokUser(config, codeChallenge)
	return sendTikTokOAuthRequest(config, code, codeVerifier)
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
	// Compute SHA-256 hash
	hash := sha256.New()
	hash.Write([]byte(codeVerifier))
	hashBytes := hash.Sum(nil)

	// Convert hash bytes to a raw hexadecimal string
	return fmt.Sprintf("%x", hashBytes)
}

func authenticateTikTokUser(config config.Config, codeChallenge string) string {
	loginRequest := request.BuildTikTokAuthorizationRequest(config, codeChallenge)
	log.Print("Invoking TikTok Login request")
	fmt.Println(loginRequest)
	_, err := client.SendRequest(loginRequest)
	if err != nil {
		panic(err)
	}
	codeChan := make(chan string)
	serverDone := &sync.WaitGroup{}
	serverDone.Add(1)
	server.StartCallbackServer(serverDone, codeChan)
	value := <-codeChan
	return value
}

func sendTikTokOAuthRequest(config config.Config, code string, codeVerifier string) model.TikTokOAuthResponse {
	loginRequest := request.BuildTikTokOAuthRequest(config, code, codeVerifier)
	log.Print("Invoking TikTok Login request")
	fmt.Println(loginRequest)
	responseBody, err := client.SendRequest(loginRequest)
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
