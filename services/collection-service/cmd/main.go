// Package main TODO(@vinjar): add proper documentation.
package main

import (
	"log/slog"
	"os"

	"innoveria-iot/collection-service/internal/server"

	"innoveria-iot/pkg/logger"
)

func main() {
	logger.NewLogger("collection-service")

	if err := server.Run(); err != nil {
		slog.Error("collection-service exited with error", "err", err)
		os.Exit(1)
	}
}
