package authctx_test

import (
	"net/http"
	"testing"

	"innoveria-iot/pkg/authctx"
)

// TestFromRequest is the test for FromRequest
func TestFromRequest(t *testing.T) {
	tests := []struct {
		name      string
		userID    string
		companyID string
		wantErr   bool
	}{
		{
			name:      "both headers present",
			userID:    "user-123",
			companyID: "company-456",
			wantErr:   false,
		},
		{
			name:      "missing user ID header",
			userID:    "",
			companyID: "company-456",
			wantErr:   true,
		},
		{
			name:      "missing company ID header",
			userID:    "user-123",
			companyID: "",
			wantErr:   true,
		},
		{
			name:      "both headers missing",
			userID:    "",
			companyID: "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest(http.MethodGet, "/", nil)

			if tt.userID != "" {
				r.Header.Set("X-Auth-User-Id", tt.userID)
			}
			if tt.companyID != "" {
				r.Header.Set("X-Auth-Company-Id", tt.companyID)
			}

			auth, err := authctx.FromRequest(r)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if auth.UserID != tt.userID {
				t.Errorf("UserID: got %q, want %q", auth.UserID, tt.userID)
			}
			if auth.CompanyID != tt.companyID {
				t.Errorf("CompanyID: got %q, want %q", auth.CompanyID, tt.companyID)
			}
		})
	}
}
