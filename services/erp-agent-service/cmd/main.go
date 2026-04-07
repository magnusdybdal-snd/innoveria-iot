// package main is the main package and entrypoint for the server
package main

import (
	"log/slog"
	"os"

	"innoveria-iot/erp-agent-service/internal/config"
	erpserviceclient "innoveria-iot/erp-agent-service/internal/erp-service-client"
	"innoveria-iot/erp-agent-service/internal/monitor"
	"innoveria-iot/erp-agent-service/internal/service"
	"innoveria-iot/pkg/logger"
)

// @title       ERP Agent Service API
// @version     1.0
// @description An edge service that fetches data from Monitor ERP for external use.

// @host        localhost:8088
func main() {
	logger.NewLogger("erp-agent-service")
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load enviroment variables", "err", err)
		os.Exit(1)
	}
	erpClient := erpserviceclient.New(cfg.ErpSvcURL)
	monitorClient := monitor.New(*cfg)

}
