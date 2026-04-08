// package main is the main package and entrypoint for the server
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"innoveria-iot/erp-agent-service/internal/config"
	erpserviceclient "innoveria-iot/erp-agent-service/internal/erp-service-client"
	"innoveria-iot/erp-agent-service/internal/monitor"
	"innoveria-iot/erp-agent-service/internal/service"
	"innoveria-iot/erp-agent-service/internal/worker"
	"innoveria-iot/pkg/logger"
)

// @title       ERP Agent Service API
// @version     1.0
// @description An edge service that fetches data from Monitor ERP for external use.

// @host        localhost:8088
func main() {
	logger.NewLogger("erp-agent-service")

	// Loading config
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load enviroment variables", "err", err)
		os.Exit(1)
	}

	// setting up the clients
	erpClient := erpserviceclient.New(cfg.ErpSvcURL)
	monitorClient := monitor.New(*cfg)

	// starting the runner, runs per cycle
	runner := service.New(monitorClient, erpClient)

	// Starting the worker loop
	worker := worker.New(cfg.PollingInterval, cfg.CycleTimeout, cfg.MaxBackoffTime, runner)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	slog.Info("starting up worker",
		"polling interval", cfg.PollingInterval,
		"Cycle timeout", cfg.CycleTimeout,
		"Max backoff time", cfg.MaxBackoffTime,
	)
	worker.Start(ctx)
}
