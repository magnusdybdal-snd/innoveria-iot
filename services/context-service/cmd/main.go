// package main is the main package and entrypoint for the server
package main

import (
	"innoveria-iot/context-service/internal/server"
	"innoveria-iot/pkg/logger"
	"log/slog"
	"os"
)

func main() {
	logger.NewLogger("context-service")

	if err := server.Run(); err != nil {
		slog.Error("context-service exited with an error", "error", err)
		os.Exit(1)
	}
}
