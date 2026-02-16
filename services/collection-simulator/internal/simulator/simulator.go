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

func New(client *broker.Client, cfg config.Config) *Simulator {
	return &Simulator{
		client:      client,
		interval:    cfg.Interval,
		topic:       cfg.Topic,
		deviceCount: cfg.Devices,
		qos:         cfg.Qos,
	}
}

func (s *Simulator) Start() {
	for i := 0; i < s.deviceCount; i++ {
		deviceId := fmt.Sprintf("%d", i)
		go s.run(deviceId)
	}

	select {}
}

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
