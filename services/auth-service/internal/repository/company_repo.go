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
	// Ordered by first to last
	findAllCompaniesQuery = `
		SELECT company_id, name, address, created_at, updated_at
		FROM auth.company
		ORDER BY created_at ASC
	`
	deleteCompanyByIDQuery = `
		DELETE FROM auth.company
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
		&out.ID,
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
	rows, err := r.db.Pool.Query(ctx, findAllCompaniesQuery)
	if err != nil {
		return nil, fmt.Errorf("find all companies: %w", err)
	}
	defer rows.Close()

	var out []domain.Company

	for rows.Next() {
		var c domain.Company
		if err := rows.Scan(&c.ID, &c.Name, &c.Address, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan company: %w", err)
		}
		out = append(out, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return out, nil
}

// FindByID retreieves all companies by id
func (r *CompanyRepoImpl) FindByID(ctx context.Context, companyID string) (domain.Company, error) {
	var out domain.Company
	err := r.db.Pool.QueryRow(ctx, findCompanyByIDQuery, companyID).Scan(
		&out.ID,
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

// DeleteByID deletes a company by id.
func (r *CompanyRepoImpl) DeleteByID(ctx context.Context, companyID string) error {
	result, err := r.db.Pool.Exec(ctx, deleteCompanyByIDQuery, companyID)
	if err != nil {
		return fmt.Errorf("delete company by id: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("delete company by id: company not found")
	}

	return nil
}
