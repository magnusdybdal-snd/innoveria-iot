package domain

import "context"

// Reconcile runs one sync pass from raw ingest data to ERP tables.
// Order should be: production resources, orders, operations, then reports.
type Reconcile interface {
	RunOnce(ctx context.Context) error
}

// ReconcileRepo defines persistence operations used by reconcile services.
// Each method processes one entity type and returns how many rows were synced.
type ReconcileRepo interface {
	SyncProductionResources(ctx context.Context, batchSize int) (int64, error)
	SyncOrders(ctx context.Context, batchSize int) (int64, error)
	SyncOrderOperations(ctx context.Context, batchSize int) (int64, error)
	SyncOrderReports(ctx context.Context, batchSize int) (int64, error)
}
