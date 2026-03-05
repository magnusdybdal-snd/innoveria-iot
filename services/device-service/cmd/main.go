// Package main TODO(@vinjar): add proper documentation.
package main

import (
	"log/slog"
	"os"

	"innoveria-iot/device-service/internal/server"
	"innoveria-iot/pkg/logger"
)

func main() {
	logger.NewLogger("device-service")

	if err := server.Run(); err != nil {
		slog.Error("device-service exited with error", "error", err)
		os.Exit(1)
	}
}
