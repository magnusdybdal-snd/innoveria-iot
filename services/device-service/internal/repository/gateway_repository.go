package repository

import (
	"context"
	"errors"
	"fmt"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgx/v5"
)

const (
	createGatewayQuery = `
		INSERT INTO device.gateway (company_id, gateway_eui, name, description, state, factory_id, factory_area_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING gateway_id, company_id, gateway_eui, name, description, state, factory_id, factory_area_id, created_at, updated_at
	`

	findGatewayByIDQuery = `
		SELECT gateway_id, company_id, gateway_eui, name, description, state, factory_id, factory_area_id, created_at, updated_at
		FROM device.gateway
		WHERE company_id = $1 AND gateway_id = $2
	`

	findAllGatewaysByCompanyIDQuery = `
		SELECT gateway_id, company_id, gateway_eui, name, description, state, factory_id, factory_area_id, created_at, updated_at
		FROM device.gateway
		WHERE company_id = $1
		ORDER BY created_at ASC
	`

	findGatewayByEUIQuery = `
		SELECT gateway_id, company_id, gateway_eui, name, description, state, factory_id, factory_area_id, created_at, updated_at
		FROM device.gateway
		WHERE gateway_eui = $1
	`

	updateGatewayStateQuery = `
		UPDATE device.gateway
		SET state = $3, updated_at = now()
		WHERE company_id = $1 AND gateway_id = $2
	`

	updateGatewayQuery = `
		UPDATE device.gateway
		SET name = $3, description = $4, factory_id = $5, factory_area_id = $6, updated_at = now()
		WHERE company_id = $1 AND gateway_id = $2
	`

	deleteGatewayQuery = `
		DELETE FROM device.gateway
		WHERE company_id = $1 AND gateway_id = $2
	`
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

	var out domain.Gateway
	err := r.db.Pool.QueryRow(ctx, createGatewayQuery,
		gateway.CompanyId,
		gateway.GatewayEUI,
		gateway.Name,
		gateway.Description,
		gateway.State,
		gateway.FactoryID,
		gateway.FactoryAreaID,
	).Scan(
		&out.Id,
		&out.CompanyId,
		&out.GatewayEUI,
		&out.Name,
		&out.Description,
		&out.State,
		&out.FactoryID,
		&out.FactoryAreaID,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		return domain.Gateway{}, fmt.Errorf("create gateway: %w", err)
	}

	return out, nil
}

// FindByID retrieves a gateway by its internal UUID.
func (r *GatewayRepository) FindByID(ctx context.Context, companyID string, gatewayID string) (domain.Gateway, error) {

	var out domain.Gateway
	err := r.db.Pool.QueryRow(ctx, findGatewayByIDQuery, companyID, gatewayID).Scan(
		&out.Id,
		&out.CompanyId,
		&out.GatewayEUI,
		&out.Name,
		&out.Description,
		&out.State,
		&out.FactoryID,
		&out.FactoryAreaID,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Gateway{}, domain.ErrNotFound
		}
		return domain.Gateway{}, fmt.Errorf("find gateway by gateway id: %w", err)
	}

	return out, nil
}

// FindAllByCompanyID retrieves all gateways belonging to the given company, ordered by creation date.
func (r *GatewayRepository) FindAllByCompanyID(ctx context.Context, companyID string) ([]domain.Gateway, error) {

	// Query the database to collect all rows
	rows, err := r.db.Pool.Query(ctx, findAllGatewaysByCompanyIDQuery, companyID)
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
			&gateway.FactoryID,
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

	var out domain.Gateway
	err := r.db.Pool.QueryRow(ctx, findGatewayByEUIQuery, gatewayEUI).Scan(
		&out.Id,
		&out.CompanyId,
		&out.GatewayEUI,
		&out.Name,
		&out.Description,
		&out.State,
		&out.FactoryID,
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
func (r *GatewayRepository) UpdateState(ctx context.Context, companyID string, gatewayID string, state domain.DeviceState) error {

	tag, err := r.db.Pool.Exec(ctx, updateGatewayStateQuery, companyID, gatewayID, state)
	if err != nil {
		return fmt.Errorf("update gateway state: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("gateway not found: %s", gatewayID)
	}

	return nil
}

// Update updates the user editable db fields of a gateway and updates the updated at timestamp.
// Returns an error if no gateway with the given ID exists.
func (r *GatewayRepository) Update(ctx context.Context, companyID string, gatewayID string, payload domain.Gateway) error {

	tag, err := r.db.Pool.Exec(ctx, updateGatewayQuery, companyID, gatewayID, payload.Name, payload.Description, payload.FactoryID, payload.FactoryAreaID)
	if err != nil {
		return fmt.Errorf("update gateway: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("update gateway: %w", domain.ErrNotFound)
	}

	return nil
}

// Delete tries to delete a gateway from the database.
// Returns an error if deletion fails or no gateway is found.
func (r *GatewayRepository) Delete(ctx context.Context, companyID string, gatewayID string) error {

	tag, err := r.db.Pool.Exec(ctx, deleteGatewayQuery, companyID, gatewayID)
	if err != nil {
		return fmt.Errorf("delete gateway %s: %w", gatewayID, err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("delete gateway: %w", domain.ErrNotFound)
	}

	return nil
}
