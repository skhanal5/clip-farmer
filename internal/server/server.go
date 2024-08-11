// Package server contains implementation details for a local HTTP server. This is only used during the authentication
// flow with TikTok's API.
package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"
)

var ctxShutdown, cancel = context.WithCancel(context.Background())

// StartCallbackServer will start a callback server at localhost:8080
// asynchronously until it receives a request at the /callback endpoint.
// This is used to handle the callback redirect that occurs when authenticating
// a TikTok user via the API.
func StartCallbackServer(wg *sync.WaitGroup, codeCh chan<- string) {
	mux := http.NewServeMux()

	// the address must match the same address you specify in
	// your TikTok developer app's callback address.
	server := http.Server{Addr: "localhost:8080", Handler: mux}

	// the endpoint must match the same endpoint you specify
	// in your TikTok developer app's callback endpoint.
	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		// look for the shutdown signal to shut down the server
		select {
		case <-ctxShutdown.Done():
			log.Println("Shutting down callback server")
			return
		default:
		}

		// TikTok should send a redirect request to this endpoint with a code representing the authentication_code and
		// some scopes. This check is to verify that this behavior occurred.
		code := r.URL.Query().Get("code")
		if code == "" {
			panic(errors.New("code and/or scope not received, authorization failed"))
		} else {
			codeCh <- code // passing the code in the channel
			cancel()
			err := server.Shutdown(context.Background()) //triggering the shutdown signal manually
			if err != nil {
				panic(err)
			}
		}
	})

	log.Print("Starting the local server")
	// handle running the server in a separate goroutine until the user completes the auth flow
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
}
