package mqtt

import (
	"encoding/json"
	"log/slog"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Collector struct {
	workers []chan ChirpstackUpEvent
}

// NewCollector starts with a buffersize and worker count
// Buffer size is the amount it can handle in a queue
// Worker count is the physical concurrent workers to read sensor data
func NewCollector(buffersize, workercount int) *Collector {
	workers := make([]chan ChirpstackUpEvent, workercount)
	for i := range workers {
		workers[i] = make(chan ChirpstackUpEvent, buffersize)
	}
	return &Collector{
		workers: workers,
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
			}
		}(i, ch)
	}
}

// Closing the queue with workers
func (c *Collector) Close() {
	for i := range c.workers {
		close(c.workers[i])
	}
}

// Hashing is used to avoid having multiple workers handle the same topic
// TODO: Test with alot of sensors to check duplication id
func hash(s string) int {
	h := 0
	for _, c := range s {
		h += int(c)
	}
	return h
}
