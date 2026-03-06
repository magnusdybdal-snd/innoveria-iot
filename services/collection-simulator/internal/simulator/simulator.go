package simulator

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"innoveria-iot/collection-simulator/internal/broker"
	"innoveria-iot/collection-simulator/internal/config"
)

type Simulator struct {
	client      *broker.Client
	interval    int
	topic       string
	deviceCount int
	qos         byte
}

// Starting a new instance of simulator
func New(client *broker.Client, cfg config.Config) *Simulator {
	return &Simulator{
		client:      client,
		interval:    cfg.Interval,
		topic:       cfg.Topic,
		deviceCount: cfg.Devices,
		qos:         cfg.Qos,
	}
}

// Initilze the simulator
// Starts goroutines so it simulate concurrent dataflow
// EUIs are generated sequentially: b000000000000001, b000000000000002, ...
// With deviceCount=3 (default) the EUIs match the dev seeds exactly.
// Increase deviceCount via the Devices env var for stress testing.
func (s *Simulator) Start() {
	for i := 0; i < s.deviceCount; i++ {
		eui := fmt.Sprintf("b%015d", i+1)
		go s.run(eui)
	}

	select {}
}

// Run is handling the data generation and mqtt publising
func (s *Simulator) run(deviceId string) {
	for {
		mock := GenerateMockEvent(deviceId)

		payload, err := json.Marshal(mock)
		if err != nil {
			slog.Error("JSON Marshal failed", "error", err)
			continue
		}
		topic := fmt.Sprintf("%s/%s", s.topic, deviceId)
		s.client.Publish(topic, s.qos, payload)

		time.Sleep(time.Duration(s.interval) * time.Second)
	}
}
