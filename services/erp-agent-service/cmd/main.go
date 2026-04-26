// Package main starts the ERP agent worker process.
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"innoveria-iot/erp-agent-service/internal/config"
	"innoveria-iot/erp-agent-service/internal/domain"
	erpserviceclient "innoveria-iot/erp-agent-service/internal/erp-service-client"
	"innoveria-iot/erp-agent-service/internal/monitor"
	"innoveria-iot/erp-agent-service/internal/monitor/mock"
	"innoveria-iot/erp-agent-service/internal/service"
	"innoveria-iot/erp-agent-service/internal/worker"
	"innoveria-iot/pkg/logger"
)

// @title       ERP Agent Service
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
	erpClient := erpserviceclient.New(cfg.ErpSvcURL, cfg.JWTToken)

	// Check enviroment to decide mock or real monitor erp
	var monitorClient domain.MonitorHandler
	if cfg.UseMockMonitor {
		monitorClient = mock.New()
		slog.Info("using mock monitor client", "GO_ENV", cfg.GOEnv)
	} else {
		monitorClient = monitor.New(*cfg)
		slog.Info("using monitor ERP client", "GO_ENV", cfg.GOEnv)
	}

	// starting the runner, runs per cycle
	runner := service.New(monitorClient, erpClient)

	// Starting the worker loop
	worker := worker.New(cfg.PollingInterval, cfg.CycleTimeout, cfg.MaxBackoffTime, runner)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	slog.Info("starting worker",
		"polling interval", cfg.PollingInterval,
		"Cycle timeout", cfg.CycleTimeout,
		"Max backoff time", cfg.MaxBackoffTime,
	)
	worker.Start(ctx)
}
