package main

import (
	"encoding/json"
	"errors"
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
		if !errors.Is(err, http.ErrServerClosed) {
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
	accept := r.Header.Get("Accept")
	var now string
	switch accept {
	case "application/json":
		w.Header().Set("Content-Type", "application/json")
		jsonTime := newJSONTime()
		out, err := json.Marshal(jsonTime)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		now = string(out)
	case "text/plain":
	default:
		w.Header().Set("Content-Type", "text/plain")
		now = time.Now().Format(time.RFC3339)
	}
	w.Write([]byte(now))
}

type JSONTime struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

func newJSONTime() JSONTime {
	now := time.Now()
	return JSONTime{
		now.Weekday().String(),
		now.Day(),
		now.Month().String(),
		now.Year(),
		now.Hour(),
		now.Minute(),
		now.Second(),
	}
}
