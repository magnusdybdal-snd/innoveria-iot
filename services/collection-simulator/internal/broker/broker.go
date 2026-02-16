package broker

import (
	"log/slog"

	"innoveria-iot/collection-simulator/internal/config"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Both sides of a mqtt networking are defined as clients
type Client struct {
	client mqtt.Client
}

func New(cfg config.Config) (*Client, error) {
	opts := mqtt.NewClientOptions().
		AddBroker(cfg.MQTTBrokerURL).
		SetClientID("simulator")

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

func (c *Client) Publish(topic string, qos byte, payload []byte) {
	token := c.client.Publish(topic, qos, false, payload)
	token.Wait()
}

func (c *Client) Close() {
	c.client.Disconnect(250)
}
