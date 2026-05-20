// package main is the main package and entrypoint for the server
package main

import (
	"innoveria-iot/context-service/internal/server"
	"innoveria-iot/pkg/logger"
	"log/slog"
	"os"
)

// @title			Context Service API
// @version			1.0
// @description 	Handles context management for sensor data

// @host 			localhost:8086
// @BasePath		/api/v1/context

func main() {
	logger.NewLogger("context-service")

	if err := server.Run(); err != nil {
		slog.Error("context-service exited with an error", "error", err)
		os.Exit(1)
	}
}
