package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"innoveria-iot/erp-service/internal/domain"
)

const (
	reconcileBatchSize = 500
)

// ReconcileImpl orchestrates one reconcile pass from raw to curated ERP data.
type ReconcileImpl struct {
	repo domain.ReconcileRepo

	mu     sync.Mutex
	cancel context.CancelFunc
	done   chan struct{}
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
	s.mu.Lock()
	if s.cancel != nil {
		s.mu.Unlock()
		slog.Warn("reconcile worker already running")
		return
	}

	runCtx, cancel := context.WithCancel(ctx)
	done := make(chan struct{})
	s.cancel = cancel
	s.done = done
	s.mu.Unlock()

	go func() {
		defer close(done)
		defer s.clearIfCurrent(done)
		// Ticker lives for the lifetime of this background worker.
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			// Stop worker on service shutdown.
			case <-runCtx.Done():
				return
			// Run one reconcile cycle per tick.
			case <-ticker.C:
				if err := s.RunOnce(runCtx); err != nil {
					if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
						return
					}
					slog.Error("reconcile cycle failed", "err", err)
				}
			}
		}
	}()
}

func (s *ReconcileImpl) clearIfCurrent(done chan struct{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.done == done {
		s.cancel = nil
		s.done = nil
	}
}

// Stop gracefully stops the background reconcile worker if it is running.
func (s *ReconcileImpl) Stop() {
	s.mu.Lock()
	cancel := s.cancel
	done := s.done
	s.mu.Unlock()

	if cancel == nil || done == nil {
		return
	}

	// Cancel the worker context and block until the goroutine exits.
	// This ensures Start/Stop cycles do not leave a running worker behind.
	cancel()
	<-done
}
