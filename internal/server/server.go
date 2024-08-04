package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var ctxShutdown, cancel = context.WithCancel(context.Background())

func StartCallbackServer(wg *sync.WaitGroup, codeCh chan<- string) {
	mux := http.NewServeMux()
	server := http.Server{Addr: "localhost:8080", Handler: mux}

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-ctxShutdown.Done():
			fmt.Println("Sorry: Shutting down ...")
			return
		default:
		}
		code := r.URL.Query().Get("code")
		fmt.Println(code)
		if code == "" {
			panic(errors.New("code and/or scope not received, authorization failed"))
		} else {
			codeCh <- code
			cancel()
			err := server.Shutdown(context.Background())
			if err != nil {
				panic(err)
			}
		}
	})

	log.Print("Starting the local server")
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
}
