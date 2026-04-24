package repository

import (
	"context"
	"errors"
	"fmt"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgx/v5"
)

const (
	createERPAgentCredentialQuery = `
		INSERT INTO auth.erp_agent_credential (company_id, secret_hash)
		VALUES ($1, $2)
		RETURNING company_id, key_id, secret_hash, created_at, rotated_at, revoked_at
	`
	findERPAgentCredentialByCompanyIDQuery = `
		SELECT company_id, key_id, secret_hash, created_at, rotated_at, revoked_at
		FROM auth.erp_agent_credential
		WHERE company_id = $1
	`
	findERPAgentCredentialByKeyIDQuery = `
		SELECT company_id, key_id, secret_hash, created_at, rotated_at, revoked_at
		FROM auth.erp_agent_credential
		WHERE key_id = $1
	`
)

// ERPAgentCredentialRepoImpl is the PostgreSQL-backed ERP agent credential repository.
type ERPAgentCredentialRepoImpl struct {
	db *dbutil.DB
}

// NewERPAgentCredentialRepo creates a new ERP agent credential repository.
func NewERPAgentCredentialRepo(db *dbutil.DB) *ERPAgentCredentialRepoImpl {
	return &ERPAgentCredentialRepoImpl{db: db}
}

// Create inserts ERP agent credential metadata for a company.
func (r *ERPAgentCredentialRepoImpl) Create(ctx context.Context, credential domain.ERPAgentCredential) (domain.ERPAgentCredential, error) {
	var out domain.ERPAgentCredential
	err := r.db.Pool.QueryRow(
		ctx,
		createERPAgentCredentialQuery,
		credential.CompanyID,
		credential.SecretHash,
	).Scan(
		&out.CompanyID,
		&out.KeyID,
		&out.SecretHash,
		&out.CreatedAt,
		&out.RotatedAt,
		&out.RevokedAt,
	)
	if err != nil {
		return domain.ERPAgentCredential{}, fmt.Errorf("create erp agent credential: %w", err)
	}

	return out, nil
}

// FindByCompanyID fetches ERP agent credential metadata by company id.
func (r *ERPAgentCredentialRepoImpl) FindByCompanyID(ctx context.Context, companyID string) (domain.ERPAgentCredential, error) {
	var out domain.ERPAgentCredential
	err := r.db.Pool.QueryRow(
		ctx,
		findERPAgentCredentialByCompanyIDQuery,
		companyID,
	).Scan(
		&out.CompanyID,
		&out.KeyID,
		&out.SecretHash,
		&out.CreatedAt,
		&out.RotatedAt,
		&out.RevokedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ERPAgentCredential{}, fmt.Errorf("find erp agent credential by company id: %w", domain.ErrERPAgentCredentialNotFound)
		}
		return domain.ERPAgentCredential{}, fmt.Errorf("find erp agent credential by company id: %w", err)
	}

	return out, nil
}

// FindByKeyID fetches ERP agent credential metadata by key id.
func (r *ERPAgentCredentialRepoImpl) FindByKeyID(ctx context.Context, keyID string) (domain.ERPAgentCredential, error) {
	var out domain.ERPAgentCredential
	err := r.db.Pool.QueryRow(
		ctx,
		findERPAgentCredentialByKeyIDQuery,
		keyID,
	).Scan(
		&out.CompanyID,
		&out.KeyID,
		&out.SecretHash,
		&out.CreatedAt,
		&out.RotatedAt,
		&out.RevokedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ERPAgentCredential{}, fmt.Errorf("find erp agent credential by key id: %w", domain.ErrERPAgentCredentialNotFound)
		}
		return domain.ERPAgentCredential{}, fmt.Errorf("find erp agent credential by key id: %w", err)
	}

	return out, nil
}
