package service

import (
	"context"

	"innoveria-iot/erp-service/internal/domain"
)

// IngestServiceImpl contains application logic for ERP ingest use-cases.
type IngestServiceImpl struct{}

// New creates a new ingest service instance.
func NewIngestService() *IngestServiceImpl {
	return &IngestServiceImpl{}
}

func (i *IngestServiceImpl) CreateOrder(ctx context.Context, payload []domain.Order) error {
	return nil
}

// CreateOrderOperation stores one order operation ingest payload.
func (i *IngestServiceImpl) CreateOrderOperation(ctx context.Context, payload []domain.OrderOperation) error {
	return nil
}

func (i *IngestServiceImpl) CreateOrderReport(ctx context.Context, payload []domain.OrderReport) error {
	return nil
}

func (i *IngestServiceImpl) CreateProductionResource(ctx context.Context, payload []domain.ProductionResource) error {
	return nil
}
