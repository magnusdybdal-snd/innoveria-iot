package services

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log/slog"

	"innoveria-iot/auth-service/internal/domain"
)

// ERPAgentCredentialServiceImpl implements ERP agent credential use-cases.
type ERPAgentCredentialServiceImpl struct {
	erpAgentCredentialRepo domain.ERPAgentCredentialRepo
}

// NewERPAgentCredentialService creates a new ERPAgentCredentialServiceImpl instance.
func NewERPAgentCredentialService(erpAgentCredentialRepo domain.ERPAgentCredentialRepo) *ERPAgentCredentialServiceImpl {
	return &ERPAgentCredentialServiceImpl{erpAgentCredentialRepo: erpAgentCredentialRepo}
}

// CreateERPAgentCredential creates ERP agent credential metadata for a company.
func (s *ERPAgentCredentialServiceImpl) CreateERPAgentCredential(ctx context.Context, credential domain.ERPAgentCredential) (domain.ERPAgentCredential, error) {
	if credential.CompanyID == "" {
		return domain.ERPAgentCredential{}, fmt.Errorf("create erp agent credential: company_id is required")
	}

	rawSecret, err := generateERPAgentSecret()
	if err != nil {
		return domain.ERPAgentCredential{}, fmt.Errorf("create erp agent credential: generate secret: %w", err)
	}

	credential.SecretHash = hashERPAgentSecret(rawSecret)

	created, err := s.erpAgentCredentialRepo.Create(ctx, credential)
	if err != nil {
		return domain.ERPAgentCredential{}, err
	}
	created.Secret = rawSecret

	slog.Info("successfully created erp agent credential")
	return created, nil
}

// GetERPAgentCredentialByCompanyID retrieves ERP agent credential metadata by company id.
func (s *ERPAgentCredentialServiceImpl) GetERPAgentCredentialByCompanyID(ctx context.Context, companyID string) (domain.ERPAgentCredential, error) {
	credential, err := s.erpAgentCredentialRepo.FindByCompanyID(ctx, companyID)
	if err != nil {
		return domain.ERPAgentCredential{}, err
	}

	slog.Info("successfully found erp agent credential by company")
	return credential, nil
}

// GetERPAgentCredentialByKeyID retrieves ERP agent credential metadata by key id.
func (s *ERPAgentCredentialServiceImpl) GetERPAgentCredentialByKeyID(ctx context.Context, keyID string) (domain.ERPAgentCredential, error) {
	credential, err := s.erpAgentCredentialRepo.FindByKeyID(ctx, keyID)
	if err != nil {
		return domain.ERPAgentCredential{}, err
	}

	slog.Info("successfully found erp agent credential by key")
	return credential, nil
}

/*
	jwt secret helper funcitons
*/

func generateERPAgentSecret() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func hashERPAgentSecret(secret string) string {
	sum := sha256.Sum256([]byte(secret))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}
