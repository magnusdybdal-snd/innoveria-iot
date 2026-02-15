package broker

import (
	"innoveria-iot/collection-simulator/internal/config"
	"log/slog"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Both sides of a mqtt networking are defined as clients
type Client struct {
	client mqtt.Client
}

func New(cfg config.Config) (*Client, error) {
	opts := mqtt.NewClientOptions().AddBroker(cfg.MQTTBrokerURL)

	opts.OnConnect = func(c mqtt.Client) {
		slog.Info("MQTT Connected")
	}

	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		slog.Error("MQTT Conntection lost", "error", err)
	}

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) Close() {
	c.client.Disconnect(250)
}
