package logger

import (
	"log/slog"
	"os"
	"strings"
)

var instance *slog.Logger
var programLevel = new(slog.LevelVar) // Info by default

func init() {
	isDebug := strings.ToLower(strings.TrimSpace(os.Getenv("FLARE_DEBUG"))) == "on"
	if isDebug {
		programLevel.Set(slog.LevelDebug)
	}

	opts := PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{Level: programLevel},
	}
	handler := NewPrettyHandler(os.Stdout, opts)
	instance = slog.New(handler)
}

func GetLogger(level string) *slog.Logger {
	switch strings.ToLower(level) {
	case "debug":
		programLevel.Set(slog.LevelDebug)
	case "info":
		programLevel.Set(slog.LevelInfo)
	case "warn":
		programLevel.Set(slog.LevelWarn)
	case "error":
		programLevel.Set(slog.LevelError)
	}
	return instance
}
