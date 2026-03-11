// Package main is the entry point for the onboarding-service
package main

import (
	"log/slog"
	"os"

	"innoveria-iot/onboarding-service/internal/server"
	"innoveria-iot/pkg/logger"
)

// @title 			Onboarding Service API
// @version 		1.0
// @description 	Manages Onboarding for new companies.

// @host 			localhost:8085
// @BasePath 		/api/v1/onboarding

func main() {
	logger.NewLogger("onboarding-service")

	if err := server.Run(); err != nil {
		slog.Error("onboarding-service exited with error", "error", err)
		os.Exit(1)
	}
}
