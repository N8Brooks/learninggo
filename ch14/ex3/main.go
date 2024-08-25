package main

import (
	"errors"
	"fmt"
	"github.com/N8Brooks/learninggo/ch14/ex3/log"
	"net/http"
	"time"
)

func main() {
	handler := http.HandlerFunc(handleHello)
	server := http.Server{
		Addr:         ":8080",
		Handler:      log.Middleware(handler),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("server closed")
		} else {
			fmt.Printf("server error: %v\n", err)
		}
	}
}

func handleHello(rw http.ResponseWriter, r *http.Request) {
	log.Log(r.Context(), log.Info, "hello called")
	rw.Header().Set("Content-Type", "text/plain")
	rw.Write([]byte("hello"))
}
