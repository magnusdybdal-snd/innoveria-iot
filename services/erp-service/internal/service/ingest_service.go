package service

import "context"

// IngestServiceImpl contains application logic for ERP ingest use-cases.
type IngestServiceImpl struct {
}

// New creates a new ingest service instance.
func New() *IngestServiceImpl {
	return &IngestServiceImpl{}
}

// CreateOrderOperation stores one order operation ingest payload.
func (i *IngestServiceImpl) CreateOrderOperation(ctx context.Context) error {
	_ = i
	_ = ctx
	return nil
}
