package simulator

import (
	"math/rand"
	"time"

	"innoveria-iot/collection-simulator/internal/domain"

	"github.com/google/uuid"
)

func GenerateMockEvent(deviceEUI string) domain.ChirpstackUpEvent {
	return domain.ChirpstackUpEvent{
		DeduplicationID: uuid.NewString(),
		Time:            time.Now().UTC().Format(time.RFC3339Nano),
		DevAddr:         "01AB23CD",
		ADR:             true,
		DR:              5,
		FCnt:            20,
		FPort:           10,
		Confirmed:       false,
		RegionConfigID:  "eu868",
		DeviceInfo: domain.DeviceInfo{
			TenantID:          "NTNU-12345",
			TenantName:        "NTNU Factory",
			ApplicationID:     "NTNU-appId",
			ApplicationName:   "K-bygget",
			DeviceProfileID:   "profile-1",
			DeviceProfileName: "Milesight EM300-DI",
			DeviceName:        "sensor-" + deviceEUI,
			DevEUI:            deviceEUI,
			DeviceClass:       "A",
			Tags:              map[string]string{},
		},
		Object: domain.SensorObject{
			Temperature: 20 + rand.Float64()*10,
			Humidity:    40 + rand.Float64()*30,
			Battery:     3.5 + rand.Float64()*0.5,
		},
	}
}
