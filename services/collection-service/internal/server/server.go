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

	"innoveria-iot/collection-service/internal/config"
	"innoveria-iot/collection-service/internal/db"
	"innoveria-iot/collection-service/internal/mqtt"
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

	// Init connection to timescale db
	db, err := db.New(cfg.DB_url)
	if err != nil {
		return fmt.Errorf("db error: %v", err)
	}
	defer db.Close()

	// Starting up a new collector
	coll := mqtt.NewCollector(1000, cfg.MQTTWorkerCount)
	coll.StartWorkers()
	defer coll.Close()

	// Starting up the mqtt client
	client, err := mqtt.New(*cfg, coll.MQTTHandler)
	if err != nil {
		return fmt.Errorf("mqtt init: %v", err)
	}
	defer client.Close()

	// subscribe to the mqtt topic
	if err := client.Subscribe(cfg.MQTTTopic); err != nil {
		return fmt.Errorf("mqtt subscribe: %v", err)
	}

	// main startup function
	serverErrors := make(chan error, 1)
	go func() {
		slog.Info("collection-service listning",
			"addr", cfg.Addr,
			"topic", cfg.MQTTTopic)

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
			return fmt.Errorf("shutdown: %w", err)
		}
	}

	return nil
}
