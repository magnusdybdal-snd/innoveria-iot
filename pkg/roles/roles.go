// Package roles defines the shared authorization role type and valid role constants.
package roles

// RoleType represents an authorization role assigned to a user.
type RoleType string

const (
	// PlatformAdmin is the platform administrator role.
	PlatformAdmin RoleType = "PLATFORM_ADMIN"
	// User is the standard factory worker role.
	User RoleType = "FACTORY_WORKER"
)
