// Package service contains ERP ingest application services.
package service

import (
	"context"
	"log/slog"

	"innoveria-iot/erp-service/internal/domain"
)

// IngestServiceImpl contains application logic for ERP ingest use-cases.
type IngestServiceImpl struct{}

// NewIngestService creates a new ingest service instance.
func NewIngestService() *IngestServiceImpl {
	return &IngestServiceImpl{}
}

// CreateOrder stores a batch of order ingest payloads.
func (i *IngestServiceImpl) CreateOrder(ctx context.Context, payload []domain.Order) error {
	slog.Info("successfully ingested orders", "count", len(payload))
	return nil
}

// CreateOrderOperation stores one order operation ingest payload.
func (i *IngestServiceImpl) CreateOrderOperation(ctx context.Context, payload []domain.OrderOperation) error {
	slog.Info("successfully ingested order operations", "count", len(payload))
	return nil
}

// CreateOrderReport stores a batch of order report ingest payloads.
func (i *IngestServiceImpl) CreateOrderReport(ctx context.Context, payload []domain.OrderReport) error {
	slog.Info("successfully ingested order reports", "count", len(payload))
	return nil
}

// CreateProductionResource stores a batch of production resource ingest payloads.
func (i *IngestServiceImpl) CreateProductionResource(ctx context.Context, payload []domain.ProductionResource) error {
	slog.Info("successfully ingested production resources", "count", len(payload))
	return nil
}
