package service

import (
	"context"
	"fmt"
	"log/slog"

	"innoveria-iot/erp-service/internal/domain"
)

const reconcileBatchSize = 500

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
	resourcesSynced, err := s.repo.SyncProductionResources(ctx, reconcileBatchSize)
	if err != nil {
		return fmt.Errorf("reconcile production resources: %w", err)
	}

	ordersSynced, err := s.repo.SyncOrders(ctx, reconcileBatchSize)
	if err != nil {
		return fmt.Errorf("reconcile orders: %w", err)
	}

	operationsSynced, err := s.repo.SyncOrderOperations(ctx, reconcileBatchSize)
	if err != nil {
		return fmt.Errorf("reconcile order operations: %w", err)
	}

	reportsSynced, err := s.repo.SyncOrderReports(ctx, reconcileBatchSize)
	if err != nil {
		return fmt.Errorf("reconcile order reports: %w", err)
	}

	slog.Info("reconcile cycle completed",
		"production_resources", resourcesSynced,
		"orders", ordersSynced,
		"order_operations", operationsSynced,
		"order_reports", reportsSynced,
	)

	return nil
}
