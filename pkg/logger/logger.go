package logger

import (
	"log/slog"
	"os"
	"strings"

	"innoveria-iot/pkg/env"
)

// NewLogger configure a slog logger for a specific service
// env:
//
//	LOG_LEVEL: debug, warn, error, info (default: info)
//	LOG_FORMAT: json, text (default: json)
func NewLogger(service string) *slog.Logger {
	level := parseLevel(env.Get("LOG_LEVEL", "info"))
	format := strings.ToLower(strings.TrimSpace(env.Get("LOG_FORMAT", "json")))

	opts := &slog.HandlerOptions{Level: level}

	var h slog.Handler
	switch format {
	case "text":
		h = slog.NewTextHandler(os.Stdout, opts)
	default:
		h = slog.NewJSONHandler(os.Stdout, opts)
	}

	logger := slog.New(h)
	if service != "" {
		logger = logger.With("service", service)
	}
	slog.SetDefault(logger)
	return logger
}

func parseLevel(level string) slog.Level {
	s := strings.ToLower(strings.TrimSpace(level))
	switch s {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
