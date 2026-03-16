// package main application entryppoint
package main

import (
	"log/slog"
	"os"

	"innoveria-iot/auth-service/internal/server"
	"innoveria-iot/pkg/logger"
)

// @title       Auth Service API
// @version     1.0
// @description Manages user authenticaion, permission roles, company generation

// @host        localhost:8084
// @BasePath    /api/v1/auth
func main() {
	logger.NewLogger("auth-service")

	if err := server.Run(); err != nil {
		slog.Error("auth-service exited with error", "error", err)
		os.Exit(1)
	}
}
