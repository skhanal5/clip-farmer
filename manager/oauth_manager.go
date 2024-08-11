package manager

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/server"
	"github.com/skhanal5/clip-farmer/internal/tiktok"
	"log"
	"math/big"
	"os"
	"strings"
	"sync"
)

func FetchTiktokOAuth(clientKey string, clientSecret string) string {
	codeVerifier := generateCodeVerifier(64)
	codeChallenge := generateRawCodeChallenge(codeVerifier)
	code := authenticateTikTokUser(clientKey, codeChallenge)
	oauth := sendTikTokOAuthRequest(clientKey, clientSecret, code, codeVerifier)
	writeToFile(oauth)
	return oauth.AccessToken
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

func authenticateTikTokUser(clientKey string, codeChallenge string) string {
	authRequest := tiktok.BuildAuthenticationRequest(clientKey, codeChallenge)
	log.Print("Invoking TikTok Login request")
	fmt.Println("Authenticate using this link: " + authRequest.URL.String())
	codeChan := make(chan string)
	serverDone := &sync.WaitGroup{}
	serverDone.Add(1)
	server.StartCallbackServer(serverDone, codeChan)
	value := <-codeChan
	return value
}

func sendTikTokOAuthRequest(clientKey string, clientSecret string, code string, codeVerifier string) tiktok.OAuthToken {
	oauthReq := tiktok.BuildOAuthRequest(clientKey, clientSecret, code, codeVerifier)
	log.Print("Invoking TikTok OAuth request")
	responseBody, err := client.SendRequest(oauthReq)
	if err != nil {
		panic(err)
	}
	var oauthResponse tiktok.OAuthToken
	err = json.Unmarshal(responseBody, &oauthResponse)
	if err != nil {
		panic(err)
	}
	log.Print("Received TikTok OAuth details")
	return oauthResponse
}

func writeToFile(o tiktok.OAuthToken) {
	file, err := os.Create("tiktok_oauth_resp.json")
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	data, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		log.Print(err)
	}
	_, err = file.Write(data)
	if err != nil {
		log.Print(err)
	}
}
