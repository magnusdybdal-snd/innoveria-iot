// Package main is the entry point for the device-service
package main

import (
	"log/slog"
	"os"

	"innoveria-iot/device-service/internal/server"
	"innoveria-iot/pkg/logger"
)

// @title 			Device Service API
// @version 		1.0
// @description 	Manages sensors, gateways, sensor profiles and sensor groups.

// @host 			localhost:8081
// @BasePath 		/api/v1/device

func main() {
	logger.NewLogger("device-service")

	if err := server.Run(); err != nil {
		slog.Error("device-service exited with error", "error", err)
		os.Exit(1)
	}
}
