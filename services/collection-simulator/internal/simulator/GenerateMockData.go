package simulator

import (
	"math/rand"
	"time"

	"innoveria-iot/collection-simulator/internal/domain"

	"github.com/google/uuid"
)

// Generate a mock sensor data
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
			TenantID:          "d0000000-0000-0000-0000-000000000001",
			TenantName:        "Innoveria Dev",
			ApplicationID:     "e0000000-0000-0000-0000-000000000001",
			ApplicationName:   "Innoveria Dev App",
			DeviceProfileID:   "f0000000-0000-0000-0000-000000000001",
			DeviceProfileName: "Dev Sensor Profile",
			DeviceName:        "sensor-" + deviceEUI,
			DevEUI:            deviceEUI,
			DeviceClass:       "A",
			Tags:              map[string]string{},
		},
		Object: buildSensorObject(deviceEUI),
	}
}

func buildSensorObject(deviceEUI string) domain.SensorObject {
	obj := domain.SensorObject{Battery: 3.5 + rand.Float64()*0.5}
	if deviceEUI == "b000000000000004" {
		obj.ElectricCurrent = 2.0 + rand.Float64()*8.0 // 2–10 A
	} else {
		obj.Temperature = 20 + rand.Float64()*10
		obj.Humidity = 40 + rand.Float64()*30
	}
	return obj
}
