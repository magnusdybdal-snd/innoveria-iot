package authctx_test

import (
	"net/http"
	"testing"

	"innoveria-iot/pkg/authctx"
)

// TestFromRequest tests that FromRequest correctly extracts all three headers.
func TestFromRequest(t *testing.T) {
	tests := []struct {
		name      string
		userID    string
		companyID string
		role      string
		wantErr   bool
	}{
		{
			name:      "all headers present",
			userID:    "user-123",
			companyID: "company-456",
			role:      authctx.RolePlatformAdmin,
			wantErr:   false,
		},
		{
			name:      "missing user ID header",
			userID:    "",
			companyID: "company-456",
			role:      authctx.RolePlatformAdmin,
			wantErr:   true,
		},
		{
			name:      "missing company ID header",
			userID:    "user-123",
			companyID: "",
			role:      authctx.RolePlatformAdmin,
			wantErr:   true,
		},
		{
			name:      "missing role header",
			userID:    "user-123",
			companyID: "company-456",
			role:      "",
			wantErr:   true,
		},
		{
			name:      "all headers missing",
			userID:    "",
			companyID: "",
			role:      "",
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
			if tt.role != "" {
				r.Header.Set("X-Auth-Role", tt.role)
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
			if auth.Role != tt.role {
				t.Errorf("Role: got %q, want %q", auth.Role, tt.role)
			}
		})
	}
}

// TestIsAdmin tests that IsAdmin correctly identifies the PLATFORM_ADMIN role.
func TestIsAdmin(t *testing.T) {
	tests := []struct {
		name string
		role string
		want bool
	}{
		{
			name: "platform admin returns true",
			role: authctx.RolePlatformAdmin,
			want: true,
		},
		{
			name: "user returns false",
			role: authctx.RoleUser,
			want: false,
		},
		{
			name: "empty role returns false",
			role: "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := authctx.Auth{Role: tt.role}
			if got := auth.IsAdmin(); got != tt.want {
				t.Errorf("IsAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}
