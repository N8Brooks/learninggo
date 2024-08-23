package main

import (
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	options := &slog.HandlerOptions{}
	handler := slog.NewJSONHandler(os.Stderr, options)
	mySlog := slog.New(handler)
	mux := http.NewServeMux()
	mux.Handle("/time", IPLogger(http.HandlerFunc(handleTime), mySlog))
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func IPLogger(h http.Handler, l *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteAddr := r.RemoteAddr
		i := strings.LastIndex(remoteAddr, ":")
		ip := remoteAddr[:i]
		l.Info("user ip", "ip", ip)
		h.ServeHTTP(w, r)
	})
}

func handleTime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	now := time.Now().Format(time.RFC3339)
	w.Write([]byte(now))
}
