// Package mqtt handles MQTT message ingestion from a Chirpstack broker.
package mqtt

import (
	"context"
	"encoding/json"
	"innoveria-iot/collection-service/internal/domain"
	"log/slog"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Collector distributes incoming MQTT uplink events across a pool of worker goroutines.
type Collector struct {
	workers []chan ChirpstackUpEvent
	service domain.MeasurementService
}

// NewCollector starts with a buffersize and worker count
// Buffer size is the amount it can handle in a queue
// Worker count is the physical concurrent workers to read sensor data
func NewCollector(buffersize, workercount int, svc domain.MeasurementService) *Collector {
	workers := make([]chan ChirpstackUpEvent, workercount)
	for i := range workers {
		workers[i] = make(chan ChirpstackUpEvent, buffersize)
	}
	return &Collector{
		workers: workers,
		service: svc,
	}
}

// MQTTHandler is how the client reacts to payload sent from a mqtt broker
func (c *Collector) MQTTHandler(client mqtt.Client, msg mqtt.Message) {
	var event ChirpstackUpEvent

	if err := json.Unmarshal(msg.Payload(), &event); err != nil {
		slog.Error("invalid uplink", "error", err)
		return
	}

	workerindex := hash(event.DeduplicationID) % len(c.workers)

	select {
	case c.workers[workerindex] <- event:
	default:
		slog.Warn("MQTT collection is full")
	}
}

// StartWorkers will start up the goroutines based on workercount
// see config.go for setting workers
func (c *Collector) StartWorkers() {
	for i, ch := range c.workers {
		go func(id int, queue chan ChirpstackUpEvent) {
			for event := range queue {
				slog.Info("processing",
					"workerId", id,
					"duplicationId", event.DeduplicationID,
					"device", event.DeviceInfo.DevEUI,
				)

				// Parse time sent by chirpstack to time.Time
				t, err := time.Parse(time.RFC3339Nano, event.Time)
				if err != nil {
					slog.Error("invalid timestamp", "device", event.DeviceInfo.DevEUI, "err", err)
					continue
				}

				payload := domain.SensorMeasurement{
					DeviceEUI: event.DeviceInfo.DevEUI,
					Timestamp: t,
					Payload:   event.Object,
				}
				if err := c.service.Create(context.Background(), payload, event.DeviceInfo.TenantID); err != nil {
					slog.Error("Failed to insert", "err", err)
				}
			}
		}(i, ch)
	}
}

// Close is closing the queue with workers
func (c *Collector) Close() {
	for i := range c.workers {
		close(c.workers[i])
	}
}

// Hashing is used to avoid having multiple workers handle the same topic
func hash(s string) int {
	h := 0
	for _, c := range s {
		h += int(c)
	}
	return h
}
