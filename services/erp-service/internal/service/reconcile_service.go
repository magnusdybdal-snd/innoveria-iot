package service

import (
	"context"
	"innoveria-iot/erp-service/internal/domain"
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
	return nil
}
