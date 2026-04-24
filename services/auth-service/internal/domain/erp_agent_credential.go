package domain

import (
	"context"
	"time"
)

// ERPAgentCredential stores per-company credential metadata for ERP agent auth.
type ERPAgentCredential struct {
	CompanyID  string
	KeyID      string
	SecretHash string
	CreatedAt  time.Time
	RotatedAt  *time.Time
	RevokedAt  *time.Time
}

// ERPAgentCredentialRepo defines persistence operations for ERP agent credentials.
type ERPAgentCredentialRepo interface {
	Create(ctx context.Context, credential ERPAgentCredential) (ERPAgentCredential, error)
	FindByCompanyID(ctx context.Context, companyID string) (ERPAgentCredential, error)
	FindByKeyID(ctx context.Context, keyID string) (ERPAgentCredential, error)
}
