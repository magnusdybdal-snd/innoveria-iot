package server

import (
	"context"
	"errors"
	"fmt"
	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/config"
	"innoveria-iot/device-service/internal/db"
	"innoveria-iot/device-service/internal/service"
	"innoveria-iot/pkg/dbutil"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() error {
	cfg := config.Load()

	// Init connection to the database
	database, err := dbutil.New(cfg.DB_url, "Device DB")
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}
	defer database.Close()

	// Setup database schema by running migrations and seeds
	if err := db.RunMigrations(database.Pool); err != nil {
		return fmt.Errorf("migrations: %w", err)
	}

	if err := db.RunSeeds(database.Pool); err != nil {
		return fmt.Errorf("seeds: %w", err)
	}

	chirpstackClient := chirpstackrest.New(*cfg)
	gatewaySvc := service.NewGatewayService(chirpstackClient)
	sensorSvc := service.NewSensorService(chirpstackClient)

	mux := NewRouter(gatewaySvc, sensorSvc)
	server := &http.Server{
		Addr:              cfg.Addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
	}
	// main startup function
	serverErrors := make(chan error, 1)
	go func() {
		slog.Info("device-service listning")
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
		slog.Info("device-service shutting down", "signal", sig.String())
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			slog.Error("device-service shutdown error", "err", err)
			err := server.Close()
			return fmt.Errorf("shutdown: %w", err)
		}
	}
	return nil
}
