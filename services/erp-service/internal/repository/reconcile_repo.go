package repository

import (
	"context"
	"embed"
	"fmt"

	"innoveria-iot/erp-service/internal/domain"
	"innoveria-iot/pkg/dbutil"
)

//go:embed sql/*.sql
var reconcileSQL embed.FS

var (
	syncProductionResourcesQuery = mustSQL("sql/sync_production_resources.sql")
	syncOrdersQuery              = mustSQL("sql/sync_orders.sql")
	syncOrderOperationsQuery     = mustSQL("sql/sync_order_operations.sql")
	syncOrderReportsQuery        = mustSQL("sql/sync_order_reports.sql")
)

func mustSQL(path string) string {
	b, err := reconcileSQL.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("load reconcile SQL %q: %v", path, err))
	}
	return string(b)
}

// ReconcileRepoImpl implements reconcile persistence operations.
type ReconcileRepoImpl struct {
	db *dbutil.DB
}

// NewReconcileRepo creates a new reconcile repository.
func NewReconcileRepo(db *dbutil.DB) *ReconcileRepoImpl {
	return &ReconcileRepoImpl{db: db}
}

func (r *ReconcileRepoImpl) execSync(ctx context.Context, query, op string, batchSize int) (int64, error) {
	res, err := r.db.Pool.Exec(ctx, query, batchSize)
	if err != nil {
		return 0, fmt.Errorf("%s: %w: %w", op, domain.ErrDatabase, err)
	}
	return res.RowsAffected(), nil
}

// SyncProductionResources syncs raw production resources into curated tables.
func (r *ReconcileRepoImpl) SyncProductionResources(ctx context.Context, batchSize int) (int64, error) {
	return r.execSync(ctx, syncProductionResourcesQuery, "sync production resources", batchSize)
}

// SyncOrders syncs raw orders into curated tables.
func (r *ReconcileRepoImpl) SyncOrders(ctx context.Context, batchSize int) (int64, error) {
	return r.execSync(ctx, syncOrdersQuery, "sync orders", batchSize)
}

// SyncOrderOperations syncs raw order operations into curated tables.
func (r *ReconcileRepoImpl) SyncOrderOperations(ctx context.Context, batchSize int) (int64, error) {
	return r.execSync(ctx, syncOrderOperationsQuery, "sync order operations", batchSize)
}

// SyncOrderReports syncs raw order reports into curated tables.
func (r *ReconcileRepoImpl) SyncOrderReports(ctx context.Context, batchSize int) (int64, error) {
	return r.execSync(ctx, syncOrderReportsQuery, "sync order reports", batchSize)
}
