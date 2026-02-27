package mappers

import (
	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
)

// MapChirpstackDeviceProfilesToDomain converts chirpstack model to domain model
func MapChirpstackDeviceProfilesToDomain(from chirpstackrest.DeviceProfile) domain.SensorProfile {
	return domain.SensorProfile{
		Id:         from.ID, // using chirpstack as of now
		Name:       from.Name,
		Region:     from.Region,
		MACVersion: from.MACVersion,
		VendorId:   from.VendorID,
		VendorName: from.VendorName,
	}
}
