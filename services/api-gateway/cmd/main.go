// Package main is the entry point for the api-gateway service.
package main

import (
	"log/slog"
	"os"

	"innoveria-iot/api-gateway/internal/server"
	"innoveria-iot/pkg/logger"
)

func main() {
	logger.NewLogger("api-gateway")

	if err := server.Run(); err != nil {
		slog.Error("api-gateway exited with error", "err", err)
		os.Exit(1)
	}
}
