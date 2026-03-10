package repository

import (
	"context"
	"fmt"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/pkg/dbutil"
)

const (
	createCompanyQuery = `
		INSERT INTO auth.company (name, address)
		VALUES ($1, $2)
		RETURNING company_id, name, address, created_at, updated_at
	`
	findCompanyByIDQuery = `
		SELECT company_id, name, address, created_at, update_at
		FROM auth.company
		WHERE company_id = $1
	`
)

// CompanyRepoImpl is the domain implementation of database operations on companies
type CompanyRepoImpl struct {
	db *dbutil.DB
}

// NewCompanyRepo initilise a new company repository
func NewCompanyRepo(db *dbutil.DB) *CompanyRepoImpl {
	return &CompanyRepoImpl{db: db}
}

// Create inserts a new company and generates a company id
func (r *CompanyRepoImpl) Create(ctx context.Context, company domain.Company) (domain.Company, error) {
	var out domain.Company
	err := r.db.Pool.QueryRow(ctx, createCompanyQuery,
		company.Name,
		company.Address,
	).Scan(
		&out.Id,
		&out.Name,
		&out.Address,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		return domain.Company{}, fmt.Errorf("create company: %w", err)
	}
	return out, nil
}

// FindAll retreieves all companies
func (r *CompanyRepoImpl) FindAll(ctx context.Context) ([]domain.Company, error) {

	return nil, nil
}

// FindByID retreieves all companies by id
func (r *CompanyRepoImpl) FindByID(ctx context.Context, companyID string) (domain.Company, error) {
	var out domain.Company
	err := r.db.Pool.QueryRow(ctx, findCompanyByIDQuery, companyID).Scan(
		&out.Id,
		&out.Name,
		&out.Address,
		&out.CreatedAt,
		&out.UpdatedAt,
	)
	if err != nil {
		return out, fmt.Errorf("find company by id: %w", err)
	}
	return out, nil
}
