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

	"innoveria-iot/context-service/internal/clients"
	"innoveria-iot/context-service/internal/clients/mock"
	"innoveria-iot/context-service/internal/config"
	"innoveria-iot/context-service/internal/db"
	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/context-service/internal/repository"
	"innoveria-iot/context-service/internal/services"
	"innoveria-iot/pkg/dbutil"
)

// Run starts the context service HTTP server and handles graceful shutdown.
func Run() error {
	cfg := config.Load()

	// Init connection to database
	database, err := dbutil.New(cfg.DB_URL, "context-db")
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}
	defer database.Close()

	if err := db.RunMigrations(database.Pool); err != nil {
		return fmt.Errorf("migrations: %w", err)
	}
	if err := db.RunSeeds(database.Pool); err != nil {
		return fmt.Errorf("seeds: %w", err)
	}
	client := clients.NewCollectionClient(cfg.CollectionSvcURL)

	var erpClient domain.ERPClient
	if cfg.UseMockERP {
		slog.Info("context-service: using mock ERP client")
		erpClient = mock.NewERPClient()
	} else {
		erpClient = clients.NewERPClient(cfg.ERPSvcURL)
	}

	repo := repository.NewRuleRepository(database)
	contextSvc := services.NewContextServiceImpl(client, erpClient, repo)
	ruleSvc := services.NewRuleServiceImpl(repo)

	// Setting up mux and http server
	mux := NewRouter(contextSvc, ruleSvc)
	server := &http.Server{
		Addr:              cfg.Addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
	}
	// main startup function
	serverErrors := make(chan error, 1)
	go func() {
		slog.Info("context-service listening")
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
		slog.Info("context-service shutting down", "signal", sig.String())
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			slog.Error("context-service shutdown error", "err", err)
			err := server.Close()
			return fmt.Errorf("shutdown: %w", err)
		}
	}
	return nil
}
