// Package server provides HTTP server setup and route definitions for the onboarding service.
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

	"innoveria-iot/onboarding-service/internal/clients"
	"innoveria-iot/onboarding-service/internal/config"
	"innoveria-iot/onboarding-service/internal/service"
)

// Run initialises dependencies, starts the HTTP server, and blocks until a shutdown signal is received.
func Run() error {
	cfg := config.Load()

	authClient := clients.NewAuthClient(cfg.AuthSvcURL)
	deviceClient := clients.NewDeviceClient(cfg.DeviceSvcURL)
	collectionClient := clients.NewCollectionClient(cfg.CollectionSvcURL)

	onboardingSvc := service.NewOnboardingService(authClient, deviceClient, collectionClient)

	mux := NewRouter(onboardingSvc, cfg.EnableSwagger)
	server := &http.Server{
		Addr:              cfg.Addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	serverErrors := make(chan error, 1)
	go func() {
		slog.Info("onboarding-service listening")
		serverErrors <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		if err == nil {
			return nil
		}
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return fmt.Errorf("listen: %w", err)
	case sig := <-shutdown:
		slog.Info("onboarding-service shutting down", "signal", sig.String())
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			slog.Error("onboarding-service shutdown error", "err", err)
			return fmt.Errorf("shutdown: %w", server.Close())
		}
	}
	return nil
}
