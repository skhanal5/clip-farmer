package manager

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/skhanal5/clip-farmer/internal/client"
	"github.com/skhanal5/clip-farmer/internal/config"
	"github.com/skhanal5/clip-farmer/internal/server"
	"github.com/skhanal5/clip-farmer/internal/tiktok"
	"log"
	"math/big"
	"os"
	"strings"
	"sync"
)

//func UploadVideosFromPath(config config.Config, path string) {
//	files, err := os.ReadDir(path)
//	if err != nil {
//		panic(err)
//	}
//
//	for _, file := range files {
//		fileInfo, err := file.Info()
//		if err != nil {
//			panic(err)
//		}
//		uploadVideoAsDraft(config, fileInfo)
//	}
//}

func UploadVideoAsDraft(config config.Config, video os.FileInfo) tiktok.FileUploadResponse {
	size := video.Size()
	sizeInMB := size / 1024 / 1024
	if sizeInMB > 64 {
		panic("file size too big to be uploaded in one chunk")
	}
	return sendFileUploadReq(config, size)
}

func sendFileUploadReq(config config.Config, size int64) tiktok.FileUploadResponse {
	videoUploadReq := tiktok.BuildVideoUploadRequest(config.TikTokOAuth.AccessToken, size)
	fmt.Println(videoUploadReq)
	res, err := client.SendRequest(videoUploadReq)
	if err != nil {
		panic(err)
	}
	var videoUploadRes tiktok.FileUploadResponse
	fmt.Println(string(res))
	err = json.Unmarshal(res, &videoUploadRes)
	if err != nil {
		panic(err)
	}
	return videoUploadRes
}

func FetchQueryCreatorInfo(config config.Config) tiktok.CreatorInfoResponse {
	creatorInfoReq := tiktok.BuildQueryCreatorInfoRequest(config.TikTokOAuth.AccessToken)
	fmt.Println(creatorInfoReq)
	res, err := client.SendRequest(creatorInfoReq)
	if err != nil {
		panic(err)
	}
	var creatorInfoRes tiktok.CreatorInfoResponse
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
	authRequest := tiktok.BuildAuthenticationRequest(config.TwitchClientId, codeChallenge)
	log.Print("Invoking TikTok Login request")
	fmt.Println(authRequest.URL.String())
	codeChan := make(chan string)
	serverDone := &sync.WaitGroup{}
	serverDone.Add(1)
	server.StartCallbackServer(serverDone, codeChan)
	value := <-codeChan
	return value
}

func sendTikTokOAuthRequest(config config.Config, code string, codeVerifier string) tiktok.OAuthResponse {
	oauthReq := tiktok.BuildOAuthRequest(config.TikTokClientKey, config.TikTokClientSecret, code, codeVerifier)
	log.Print("Invoking TikTok OAuth request")
	responseBody, err := client.SendRequest(oauthReq)
	if err != nil {
		panic(err)
	}
	var oauthResponse tiktok.OAuthResponse
	fmt.Println(string(responseBody))
	err = json.Unmarshal(responseBody, &oauthResponse)
	if err != nil {
		panic(err)
	}
	return oauthResponse
}
