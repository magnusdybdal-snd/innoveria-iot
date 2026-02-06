package main

import (
	"innoveria-iot/api-gateway/internal/server"
	"innoveria-iot/pkg/logger"
	"log/slog"
	"os"
)

// TODO: Setup logging
func main()  {
	logger.NewLogger("api-gateway")

	if err := server.Run(); err != nil {
		slog.Error("api-gateway exited with error", "err", err)
		os.Exit(1)
	}
}
