package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"innoveria-iot/api-gateway/internal/config"
)

// Run is the server entry point
func Run() error {
	cfg := config.Load()
	mux := NewRouter(cfg)

	server := &http.Server{
		Addr:              cfg.Addr,
		Handler:           corsMiddleware(mux),
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
		// Intentionally avoid ReadTimeout/WriteTimeout here because of mqtt
	}

	// main startup function
	serverErrors := make(chan error, 1)
	go func() {
		slog.Info("api-gateway listning", "addr", cfg.Addr)
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
		slog.Info("api-gateway shutting down", "signal", sig.String())
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			slog.Error("api-gateway shutdown error", "err", err)
			err := server.Close()
			if err != nil {
				log.Printf("Error closing server: %v\n", err)
			}

			return fmt.Errorf("shutdown: %w", err)
		}
	}

	return nil
}
