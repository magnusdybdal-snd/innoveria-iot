package main

import (
	"log/slog"
	"os"

	"innoveria-iot/auth-service/internal/server"
	"innoveria-iot/pkg/logger"
)

func main() {
	logger.NewLogger("auth-service")

	if err := server.Run(); err != nil {
		slog.Error("device-service exited with error", "error", err)
		os.Exit(1)
	}
}
