package domain

import (
	"context"
)

// Ingest defines write operations for ERP ingest payloads.
type Ingest interface {
	CreateOrder(ctx context.Context, payload []Order) error
	CreateOrderOperation(ctx context.Context, payload []OrderOperation) error
	CreateOrderReport(ctx context.Context, payload []OrderReport) error
	CreateProductionResource(ctx context.Context, payload []ProductionResource) error
}

// IngestRepo defines the write operation for persistant storage
type IngestRepo interface {
	CreateOrder(ctx context.Context, payload []Order) error
	CreateOrderOperation(ctx context.Context) error
	CreateOrderReport(ctx context.Context) error
	CreateProductionResource(ctx context.Context) error
}
