package service

import (
	"context"
	"fmt"

	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/chirpstackrest/dto"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/service/mappers"
)

// SensorProfileServiceImpl implements domain.SensorProfileService, fetching sensor profiles from Chirpstack.
type SensorProfileServiceImpl struct {
	cc *chirpstackrest.Client
}

// NewSensorProfileService creates a new SensorProfileServiceImpl with the given Chirpstack client.
func NewSensorProfileService(cc *chirpstackrest.Client) *SensorProfileServiceImpl {
	return &SensorProfileServiceImpl{
		cc: cc,
	}
}

// GetAll retrieves all available sensor profiles from Chirpstack up to the given limit.
func (s *SensorProfileServiceImpl) GetAll(ctx context.Context, limit int) ([]domain.SensorProfile, error) {
	resp, err := s.cc.GetAllSensorProfiles(ctx, limit)
	if err != nil {
		return nil, err
	}
	var result []domain.SensorProfile
	for _, sp := range resp.Result {
		result = append(result, mappers.MapChirpstackDeviceProfilesToDomain(sp))
	}

	return result, nil
}

// GetOne is not yet implemented.
func (s *SensorProfileServiceImpl) GetOne(ctx context.Context) (domain.SensorProfile, error) {
	return domain.SensorProfile{}, nil
}

// EnsureTenantProfile ensures a tenant-level copy of the given profile exists.
// profileID may point to either a global profile or an already tenant-owned profile.
// If it is already owned by the tenant it is returned as-is.
// Otherwise a tenant-level copy is created (or reused if one with the same name already exists).
// Returns the tenant-level profile ID to use when registering a device.
func (s *SensorProfileServiceImpl) EnsureTenantProfile(ctx context.Context, profileID string, tenantID string) (string, error) {
	// Fetch the profile by ID — works for both global and tenant-level profiles.
	profile, err := s.cc.GetSensorProfile(ctx, profileID)
	if err != nil {
		return "", fmt.Errorf("ensure tenant profile: fetch profile: %w", err)
	}

	// Profile is already owned by this tenant — nothing to do.
	if profile.TenantID == tenantID {
		return profile.ID, nil
	}

	// Profile is global — check if a tenant-level copy already exists by name.
	// Note: GetTenantSensorProfiles returns both global and tenant-level profiles mixed,
	// so we verify ownership of each match individually.
	existing, err := s.cc.GetTenantSensorProfiles(ctx, tenantID, 1000)
	if err != nil {
		return "", fmt.Errorf("ensure tenant profile: fetch existing: %w", err)
	}

	if existing.TotalCount > len(existing.Result) {
		return "", fmt.Errorf("ensure tenant profile: result truncated (%d of %d profiles returned) — increase limit", len(existing.Result), existing.TotalCount)
	}

	for _, p := range existing.Result {
		if p.Name == profile.Name {
			owned, err := s.cc.GetSensorProfile(ctx, p.ID)
			if err != nil {
				return "", fmt.Errorf("ensure tenant profile: verify existing profile: %w", err)
			}
			if owned.TenantID == tenantID {
				return owned.ID, nil
			}
		}
	}

	// No tenant-level copy found — create one from the global profile.
	id, err := s.cc.CreateTenantSensorProfile(ctx, dto.CreateDeviceProfileRequest{
		DeviceProfile: dto.CreateDeviceProfileBody{
			TenantID:                tenantID,
			Name:                    profile.Name,
			Description:             profile.Description,
			Region:                  profile.Region,
			MACVersion:              profile.MACVersion,
			RegParamsRevision:       profile.RegParamsRevision,
			ADRAlgorithmID:          profile.ADRAlgorithmID,
			PayloadCodecRuntime:     profile.PayloadCodecRuntime,
			PayloadCodecScript:      profile.PayloadCodecScript,
			FlushQueueOnActivate:    profile.FlushQueueOnActivate,
			UplinkInterval:          profile.UplinkInterval,
			DeviceStatusReqInterval: profile.DeviceStatusReqInterval,
			SupportsOtaa:            profile.SupportsOtaa,
			SupportsClassB:          profile.SupportsClassB,
			SupportsClassC:          profile.SupportsClassC,
			Tags:                    profile.Tags,
		},
	})
	if err != nil {
		return "", fmt.Errorf("ensure tenant profile: create tenant profile: %w", err)
	}

	return id, nil
}
