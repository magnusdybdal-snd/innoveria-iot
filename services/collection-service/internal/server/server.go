package server

import (
	"context"
	"errors"
	"fmt"
	"innoveria-iot/collection-service/internal/config"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server entry point
func Run() error {
	cfg := config.Load()
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
		slog.Info("collection-service listning", "addr", cfg.Addr)
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
		slog.Info("collection-service shutting down", "signal", sig.String())
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			slog.Error("collection-service shutdown error", "err", err)
			err := server.Close()
			if err != nil {
				log.Printf("Error closing server: %v\n", err)
			}
			return fmt.Errorf("shutdown: %w", err)
		}
	}

	return nil
}
