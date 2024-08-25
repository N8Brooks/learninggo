package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	mux := http.NewServeMux()
	timeoutMiddleware := createTimeoutMiddleware(1000)
	delayHandler := http.HandlerFunc(handleDelay)
	mux.Handle("/delay/{seconds}", delayHandler)
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      timeoutMiddleware(mux),
	}
	err := s.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server error: %v\n", err)
		}
	}
}

func createTimeoutMiddleware(ms int) func(http.Handler) http.Handler {
	timeout := time.Duration(ms) * time.Millisecond
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

func handleDelay(w http.ResponseWriter, r *http.Request) {
	seconds, err := strconv.ParseFloat(r.PathValue("seconds"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	message, err := delay(r.Context(), seconds)
	w.Header().Set("Content-Type", "text/plain")
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusRequestTimeout)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write([]byte(message))
}

func delay(ctx context.Context, seconds float64) (string, error) {
	duration := time.Duration(seconds * float64(time.Second))
	select {
	case <-time.After(duration):
		return "Done", nil
	case <-ctx.Done():
		return "Too slow", ctx.Err()
	}
}
