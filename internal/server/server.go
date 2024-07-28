package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	Code = ""
)

var ctxShutdown, cancel = context.WithCancel(context.Background())

func StartServer(wg *sync.WaitGroup) {
	mux := http.NewServeMux()
	server := http.Server{Addr: "localhost:8080", Handler: mux}

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-ctxShutdown.Done():
			fmt.Println("Sorry: Shuting down ...")
			return
		default:
		}

		fmt.Println(r)
		Code = r.URL.Query().Get("code")
		if Code == "" {
			panic(errors.New("code and/or scope not received, authorization failed"))
		} else {
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
			fmt.Println(err)
		}
	}()
}

func verifyCSRFState(stateAfterRequest string, stateBeforeRequest string) error {
	if stateAfterRequest != stateBeforeRequest {
		panic(errors.New("CSRF state does not match"))
	}
	return nil
}
