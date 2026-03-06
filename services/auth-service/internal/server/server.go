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

	"innoveria-iot/auth-service/internal/config"
	"innoveria-iot/pkg/dbutil"
)

// Run starts the auth service HTTP server and handles graceful shutdown.
func Run() error {
	cfg := config.Load()

	// Init connection to auth database
	database, err := dbutil.New(cfg.DB_URL, "auth-db")
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}
	defer database.Close()

	// Setting up mux and http server
	mux := NewRouter()
	server := &http.Server{
		Addr:              cfg.Addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
	}
	// main startup function
	serverErrors := make(chan error, 1)
	go func() {
		slog.Info("auth-service listning")
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
		slog.Info("auth-service shutting down", "signal", sig.String())
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			slog.Error("auth-service shutdown error", "err", err)
			err := server.Close()
			return fmt.Errorf("shutdown: %w", err)
		}
	}
	return nil
}
