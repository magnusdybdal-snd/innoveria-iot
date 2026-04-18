package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"innoveria-iot/erp-service/internal/domain"
)

const (
	reconcileBatchSize = 500
)

// ReconcileImpl orchestrates one reconcile pass from raw to curated ERP data.
type ReconcileImpl struct {
	repo domain.ReconcileRepo
}

// NewReconcileService creates a reconcile service with the required repository.
func NewReconcileService(repo domain.ReconcileRepo) *ReconcileImpl {
	return &ReconcileImpl{
		repo: repo,
	}
}

// RunOnce performs a single reconciliation cycle.
func (s *ReconcileImpl) RunOnce(ctx context.Context) error {
	var reconcileErr error
	resourcesSynced, err := s.repo.SyncProductionResources(ctx, reconcileBatchSize)
	if err != nil {
		reconcileErr = errors.Join(reconcileErr, err)
	}

	ordersSynced, err := s.repo.SyncOrders(ctx, reconcileBatchSize)
	if err != nil {
		reconcileErr = errors.Join(reconcileErr, err)
	}

	operationsSynced, err := s.repo.SyncOrderOperations(ctx, reconcileBatchSize)
	if err != nil {
		reconcileErr = errors.Join(reconcileErr, err)
	}

	reportsSynced, err := s.repo.SyncOrderReports(ctx, reconcileBatchSize)
	if err != nil {
		reconcileErr = errors.Join(reconcileErr, err)
	}

	slog.Info("reconcile cycle completed",
		"production_resources", resourcesSynced,
		"orders", ordersSynced,
		"order_operations", operationsSynced,
		"order_reports", reportsSynced,
	)
	if reconcileErr != nil {
		return fmt.Errorf("reconcile error: %w", reconcileErr)
	}

	return nil
}

// Start runs reconciliation on a fixed interval until ctx is cancelled.
// It starts a background goroutine and returns immediately.
// interval is check in config.go
func (s *ReconcileImpl) Start(ctx context.Context, interval time.Duration) {
	go func() {
		// Ticker lives for the lifetime of this background worker.
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			// Stop worker on service shutdown.
			case <-ctx.Done():
				return
			// Run one reconcile cycle per tick.
			case <-ticker.C:
				if err := s.RunOnce(ctx); err != nil {
					slog.Error("reconcile cycle failed", "err", err)
				}
			}
		}
	}()
}
