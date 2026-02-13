package mqtt

import (
	"encoding/json"
	"log"
	"log/slog"

	"innoveria-iot/collection-service/internal/models"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Collector struct {
	queue chan models.ChirpstackUpEvent
}

func NewCollector(buffersize int) *Collector {
	return &Collector{
		queue: make(chan models.ChirpstackUpEvent, buffersize),
	}
}

func (c *Collector) MQTTHandler(client mqtt.Client, msg mqtt.Message) {
	var event models.ChirpstackUpEvent

	if err := json.Unmarshal(msg.Payload(), &event); err != nil {
		slog.Error("invalid uplink", "error", err)
		return
	}

	select {
	case c.queue <- event:
	default:
		slog.Warn("MQTT collection is full")
	}
}

func (c *Collector) StartWorker(workercount int) {
	for i := 0; i < workercount; i++ {
		go func (id int) {
			for event := range c.queue {
				log.Printf("worker(%d): %v\n", id, event) // replaced with database handling
			}
		}(i)
	}
}

func (c *Collector) Close() {
	close(c.queue)
}
