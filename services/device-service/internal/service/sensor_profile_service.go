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

// EnsureTenantProfile ensures a tenant-level copy of the given global profile exists.
// If one already exists (matched by name), its ID is returned. Otherwise a new one is created.
// Returns the tenant-level profile ID to use when registering a device.
func (s *SensorProfileServiceImpl) EnsureTenantProfile(ctx context.Context, globalProfileID string, tenantID string) (string, error) {
	// Fetch existing tenant profiles to check for a duplicate by name.
	existing, err := s.cc.GetTenantSensorProfiles(ctx, tenantID, 1000)
	if err != nil {
		return "", fmt.Errorf("ensure tenant profile: fetch existing: %w", err)
	}

	// Fetch the global profile to get its name and full settings.
	global, err := s.cc.GetSensorProfile(ctx, globalProfileID)
	if err != nil {
		return "", fmt.Errorf("ensure tenant profile: fetch global profile: %w", err)
	}

	// Return existing tenant profile if one with the same name already exists.
	for _, p := range existing.Result {
		if p.Name == global.Name {
			return p.ID, nil
		}
	}

	// No existing tenant profile found — create one from the global profile.
	id, err := s.cc.CreateTenantSensorProfile(ctx, dto.CreateDeviceProfileRequest{
		DeviceProfile: dto.CreateDeviceProfileBody{
			TenantID:                tenantID,
			Name:                    global.Name,
			Description:             global.Description,
			Region:                  global.Region,
			MACVersion:              global.MACVersion,
			RegParamsRevision:       global.RegParamsRevision,
			ADRAlgorithmID:          global.ADRAlgorithmID,
			PayloadCodecRuntime:     global.PayloadCodecRuntime,
			PayloadCodecScript:      global.PayloadCodecScript,
			FlushQueueOnActivate:    global.FlushQueueOnActivate,
			UplinkInterval:          global.UplinkInterval,
			DeviceStatusReqInterval: global.DeviceStatusReqInterval,
			SupportsOtaa:            global.SupportsOtaa,
			SupportsClassB:          global.SupportsClassB,
			SupportsClassC:          global.SupportsClassC,
			Tags:                    global.Tags,
		},
	})
	if err != nil {
		return "", fmt.Errorf("ensure tenant profile: create tenant profile: %w", err)
	}

	return id, nil
}
