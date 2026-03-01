package repository

import (
	"context"
	"fmt"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/pkg/dbutil"
)

// GatewayRepository handles persistence of gateway metadata stored in the database.
type GatewayRepository struct {
	db *dbutil.DB
}

// NewGatewayRepository creates a new GatewayRepository with the given database.
func NewGatewayRepository(db *dbutil.DB) *GatewayRepository {
	return &GatewayRepository{db: db}
}

// Create inserts a new gateway into the database.
// Returns the newly created gateway with database generated fields.
func (r *GatewayRepository) Create(ctx context.Context, gateway domain.Gateway) (domain.Gateway, error) {
	const query = `
		INSERT INTO device.gateway (company_id, gateway_eui, name, description, state, factory_area_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING gateway_id, company_id, gateway_eui, name, description, state, factory_area_id, created_at, updated_at
	`

	var out domain.Gateway
	err := r.db.Pool.QueryRow(ctx, query,
		gateway.CompanyId,
		gateway.GatewayEUI,
		gateway.Name,
		gateway.Description,
		gateway.State,
		gateway.FactoryAreaID,
	).Scan(
		&out.Id,
		&out.CompanyId,
		&out.GatewayEUI,
		&out.Name,
		&out.Description,
		&out.State,
		&out.FactoryAreaID,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		return domain.Gateway{}, fmt.Errorf("create gateway: %w", err)
	}

	return out, nil
}

// FindByID retrieves a gateway by its internal UUID
func (r *GatewayRepository) FindByID(ctx context.Context, gatewayID string) (domain.Gateway, error) {
	const query = `
		SELECT gateway_id, company_id, gateway_eui, name, description, state, factory_area_id, created_at, updated_at
		FROM device.gateway
		WHERE gateway_id = $1
	`

	var out domain.Gateway
	err := r.db.Pool.QueryRow(ctx, query, gatewayID).Scan(
		&out.Id,
		&out.CompanyId,
		&out.GatewayEUI,
		&out.Name,
		&out.Description,
		&out.State,
		&out.FactoryAreaID,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		return domain.Gateway{}, fmt.Errorf("find gateway by gateway id: %w", err)
	}

	return out, nil
}

// FindAllByCompanyID retrieves all gateways belonging to the given company, ordered by creation date.
func (r *GatewayRepository) FindAllByCompanyID(ctx context.Context, companyID string) ([]domain.Gateway, error) {
	const query = `
		SELECT gateway_id, company_id, gateway_eui, name, description, state, factory_area_id, created_at, updated_at
		FROM device.gateway
		WHERE company_id = $1
		ORDER BY created_at ASC
	`

	// Query the database to collect all rows
	rows, err := r.db.Pool.Query(ctx, query, companyID)
	if err != nil {
		return nil, fmt.Errorf("find all gateways by company id: %w", err)
	}
	defer rows.Close()

	// The slice of gateways to be returned
	var out []domain.Gateway

	for rows.Next() {
		var gateway domain.Gateway

		err := rows.Scan(
			&gateway.Id,
			&gateway.CompanyId,
			&gateway.GatewayEUI,
			&gateway.Name,
			&gateway.Description,
			&gateway.State,
			&gateway.FactoryAreaID,
			&gateway.CreatedAt,
			&gateway.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan gateway: %w", err)
		}

		// Add the gateway to the slice
		out = append(out, gateway)
	}

	// Sanity check if the loop ended due to an error
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return out, nil

}

// FindByEUI retrieves a gateway by its hardware EUI. Used for deduplication checks before registration.
func (r *GatewayRepository) FindByEUI(ctx context.Context, gatewayEUI string) (domain.Gateway, error) {
	const query = `
		SELECT gateway_id, company_id, gateway_eui, name, description, state, factory_area_id, created_at, updated_at
		FROM device.gateway
		WHERE gateway_eui = $1
	`

	var out domain.Gateway
	err := r.db.Pool.QueryRow(ctx, query, gatewayEUI).Scan(
		&out.Id,
		&out.CompanyId,
		&out.GatewayEUI,
		&out.Name,
		&out.Description,
		&out.State,
		&out.FactoryAreaID,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		return domain.Gateway{}, fmt.Errorf("find gateway by gateway eui: %w", err)
	}

	return out, nil
}

// UpdateState sets the administrative state of a gateway and updates the updated at timestamp.
// Returns an error if no gateway with the given ID exists.
func (r *GatewayRepository) UpdateState(ctx context.Context, gatewayID string, state domain.DeviceState) error {
	const query = `
		UPDATE device.gateway
		SET state = $1, updated_at = now()
		WHERE gateway_id = $2
	`
	tag, err := r.db.Pool.Exec(ctx, query, state, gatewayID)
	if err != nil {
		return fmt.Errorf("update gateway state: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("gateway not found: %s", gatewayID)
	}

	return nil
}
