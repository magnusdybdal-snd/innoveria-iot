package main

import (
	"log/slog"

	"innoveria-iot/device-service/internal/server"
	"innoveria-iot/pkg/logger"
)

func main() {
	logger.NewLogger("collection-service")

	if err := server.Run(); err != nil {
		slog.Error("device-service exited with error", "error", err)
	}
}
