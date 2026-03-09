// Package main is the entry point for the collection-service
package main

import (
	"log/slog"
	"os"

	"innoveria-iot/collection-service/internal/server"

	"innoveria-iot/pkg/logger"
)

// @title			Collection Service API
// @version			1.0
// @description 	Handles ingestion and retrieval of sensor measurements

// @host 			localhost:8082
// @BasePath		/api/v1/collection

func main() {
	logger.NewLogger("collection-service")

	if err := server.Run(); err != nil {
		slog.Error("collection-service exited with error", "err", err)
		os.Exit(1)
	}
}
