package tiktok

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/server"
	"log"
	"math/big"
	"os"
	"strings"
	"sync"
)

// FetchTiktokOAuth navigates the TikTok authentication flow and retrieves an OAuth token in a two-stage process.
// The first stage entails authenticating into the user, we want to post on the behalf of. This is a blocking process
// until the user finishes. The second half is retrieving an OAuth token using the authentication details and returning
// it as a string.
func FetchTiktokOAuth(clientKey string, clientSecret string) string {
	codeVerifier := generateCodeVerifier(64)
	codeChallenge := generateRawCodeChallenge(codeVerifier)
	code := authenticateTikTokUser(clientKey, codeChallenge)
	oauth := sendTikTokOAuthRequest(clientKey, clientSecret, code, codeVerifier)
	writeToFile(oauth)
	return oauth.AccessToken
}

// generateCodeVerifier generates a code verifier of the specified length randomly with the designated character.
// Returns a string that is not url encoded.
func generateCodeVerifier(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~"
	var result strings.Builder
	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		result.WriteByte(chars[index.Int64()])
	}
	return result.String()
}

// generateRawCodeChallenge generates a hashed string from the codeVerifier. Returns another string
// with the hashed contents and is not url encoded.
func generateRawCodeChallenge(codeVerifier string) string {
	hash := sha256.New()
	hash.Write([]byte(codeVerifier))
	hashBytes := hash.Sum(nil)
	return fmt.Sprintf("%x", hashBytes)
}

// authenticateTikTokUser handles authenticating the user by going through TikTok's authentication flow.
// Blocks until the user authenticates manually with the link that is outputted.
// Returns a string representing the authentication_code
func authenticateTikTokUser(clientKey string, codeChallenge string) string {
	authRequest := BuildAuthenticationRequest(clientKey, codeChallenge)
	log.Print("Invoking TikTok Login request")

	// Generates the authentication URL containing all necessary scopes
	fmt.Println("Authenticate using this link: " + authRequest.URL.String())

	codeChan := make(chan string)
	serverDone := &sync.WaitGroup{}
	serverDone.Add(1)

	//starts server in the background and blocks until the user finishes the authentication flow on the browser.
	server.StartCallbackServer(serverDone, codeChan)

	value := <-codeChan
	return value
}

// sendTikTokOAuthRequest wraps up the OAuth flow by sending the OAuth request with the authentication_code and returns
// a struct representing an OAuth token
func sendTikTokOAuthRequest(clientKey string, clientSecret string, code string, codeVerifier string) OAuthToken {
	oauthReq := BuildOAuthRequest(clientKey, clientSecret, code, codeVerifier)
	log.Print("Invoking TikTok OAuth request")
	responseBody, err := client.SendRequest(oauthReq)
	if err != nil {
		panic(err)
	}
	var oauthResponse OAuthToken
	err = json.Unmarshal(responseBody, &oauthResponse)
	if err != nil {
		panic(err)
	}
	log.Print("Received TikTok OAuth details")
	return oauthResponse
}

// writeToFile writes the OAuthToken to a file in the root directory for later re-use
func writeToFile(o OAuthToken) {
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
