// Package server provides HTTP server for startup
package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"innoveria-iot/erp-service/internal/config"
	"innoveria-iot/erp-service/internal/db"
	"innoveria-iot/erp-service/internal/repository"
	"innoveria-iot/erp-service/internal/service"
	"innoveria-iot/pkg/dbutil"
)

// Run starts the erp service HTTP server and handles graceful shutdown.
func Run() error {
	cfg := config.Load()

	// Init connection to erp database
	database, err := dbutil.New(cfg.DBURL, "erp-db")
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}
	defer database.Close()

	// setup database with migrations
	if err := db.RunMigrations(database.Pool); err != nil {
		return fmt.Errorf("migrations: %w", err)
	}

	// repository init
	ingestRepo := repository.NewIngestRepo(database)
	reconcileRepo := repository.NewReconcileRepo(database)
	prodResRepo := repository.NewProductionResourceRepo(database)

	// Service init
	ingestSvc := service.NewIngestService(ingestRepo)
	prodResSvc := service.NewProductionResourceSvc(prodResRepo)

	// Reconcile worker
	workerCtx, workerCancel := context.WithCancel(context.Background())
	defer workerCancel()

	reconcileWorker := service.NewReconcileService(reconcileRepo)
	reconcileWorker.Start(workerCtx, cfg.ReconcileInterval)
	defer reconcileWorker.Stop()

	// Setting up mux and http server
	mux := NewRouter(ingestSvc, prodResSvc)
	server := &http.Server{
		Addr:              cfg.Addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
	}
	// main startup function
	serverErrors := make(chan error, 1)
	go func() {
		slog.Info("erp-service listning")
		err := server.ListenAndServe()
		serverErrors <- err
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Handles server startup errors and graceful shutdown
	select {
	case err := <-serverErrors:
		if err == nil {
			return nil
		}
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return fmt.Errorf("listen %w", err)
	case sig := <-shutdown:
		slog.Info("erp-service shutting down", "signal", sig.String())
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			slog.Error("erp-service shutdown error", "err", err)
			err := server.Close()
			return fmt.Errorf("shutdown: %w", err)
		}
	}
	return nil
}
