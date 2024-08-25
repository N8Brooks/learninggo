package log

import (
	"context"
	"fmt"
	"net/http"
)

type LogLevel string

const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
)

func Log(ctx context.Context, level LogLevel, message string) {
	inLevel, ok := GetLogLevel(ctx)
	if !ok {
		return
	}
	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}

type logLevelKey struct{}

func SetLogLevel(ctx context.Context, level LogLevel) context.Context {
	return context.WithValue(ctx, logLevelKey{}, level)
}

func GetLogLevel(ctx context.Context) (LogLevel, bool) {
	logLevel, ok := ctx.Value(logLevelKey{}).(LogLevel)
	return logLevel, ok
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if level := LogLevel(r.URL.Query().Get("log_level")); level == Debug || level == Info {
			ctx := r.Context()
			ctx = SetLogLevel(ctx, LogLevel(level))
			h.ServeHTTP(w, r.WithContext(ctx))
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
