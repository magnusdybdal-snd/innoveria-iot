// package main is the main package and entrypoint for the server
package main

import (
	"log/slog"
	"os"

	"innoveria-iot/erp-agent-service/internal/server"
	"innoveria-iot/pkg/logger"
)

// @title       ERP Agent Service API
// @version     1.0
// @description

// @host        localhost:8088
// @BasePath    /api/v1/erp-agent
func main() {
	logger.NewLogger("erp-agent-service")

	if err := server.Run(); err != nil {
		slog.Error("erp-agent-service exited with error", "error", err)
		os.Exit(1)
	}
}
