// Package main is the erp-service application entrypoint.
package main

import (
	"log/slog"
	"os"

	"innoveria-iot/erp-service/internal/server"
	"innoveria-iot/pkg/logger"
)

// @title       Erp Service API
// @version     1.0
// @description Manages production resource data from monitor erp

// @host        localhost:8087
// @BasePath    /api/v1/erp
func main() {
	logger.NewLogger("erp-service")

	if err := server.Run(); err != nil {
		slog.Error("erp-service exited with error", "error", err)
		os.Exit(1)
	}
}
