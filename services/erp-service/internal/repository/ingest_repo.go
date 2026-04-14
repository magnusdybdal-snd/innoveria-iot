package repository

import (
	"context"

	"innoveria-iot/pkg/dbutil"
)

// IngestRepoImpl is the implementaion of ingest persistant storage write
// Which is responsible for writing erp data from a erp agent service
type IngestRepoImpl struct {
	db *dbutil.DB
}

// NewIngestRepo initializes a new Ingest repository.
func NewIngestRepo(db *dbutil.DB) *IngestRepoImpl {
	return &IngestRepoImpl{
		db: db,
	}
}

func (r *IngestRepoImpl) CreateOrder(ctx context.Context) error {
	return nil
}

func (r *IngestRepoImpl) CreateOrderOperation(ctx context.Context) error {
	return nil
}

func (r *IngestRepoImpl) CreateOrderReport(ctx context.Context) error {
	return nil
}

func (r *IngestRepoImpl) CreateProductionResource(ctx context.Context) error {
	return nil
}
