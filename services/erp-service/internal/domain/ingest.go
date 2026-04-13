package domain

import (
	"context"
)

type Ingest interface {
	CreateOrder(ctx context.Context, payload []Order) error
	CreateOrderOperation(ctx context.Context, payload []OrderOperation) error
	CreateOrderReport(ctx context.Context, payload []OrderReport) error
	CreateProductionResource(ctx context.Context, payload []ProductionResource) error
}
