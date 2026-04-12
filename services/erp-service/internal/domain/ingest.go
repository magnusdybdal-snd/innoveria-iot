package domain

import "context"

type Ingest interface {
	CreateOrderOperation(ctx context.Context) error
	CreateOrderReporting(ctx context.Context) error
	CreateWorkCenter(ctx context.Context) error
}
