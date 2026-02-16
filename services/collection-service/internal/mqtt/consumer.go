package mqtt

import (
	"innoveria-iot/collection-service/internal/config"
	"log/slog"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const disconnectTime = 250 // in millis

type Client struct {
	client mqtt.Client
}

// New generates a mqtt client for listning to topics from a mqtt broker
func New(cfg config.Config, msgHandler mqtt.MessageHandler) (*Client, error) {
	opts := mqtt.NewClientOptions().
		AddBroker(cfg.MQTTBrokerURL).
		SetClientID(cfg.MQTTClientId).
		SetAutoReconnect(true).
		SetConnectRetry(false).
		SetConnectRetryInterval(5 * time.Second).
		SetKeepAlive(30 * time.Second).
		SetPingTimeout(10 * time.Second).
		SetCleanSession(false)

	// Connect success
	opts.OnConnect = func(c mqtt.Client) {
		slog.Info("MQTT Connected")
	}

	// error with connection
	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		slog.Error("MQTT Conntection lost", "error", err)
	}

	// How we structure mqtt messages
	opts.DefaultPublishHandler = msgHandler

	// Starting the client with a configuration
	client := mqtt.NewClient(opts)

	// error check with client connection
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &Client{client: client}, nil
}

// Subscribe will listen to a mqtt topic based on config.go
func (c *Client) Subscribe(topic string) error {
	token := c.client.Subscribe(topic, 0, nil) // qos is 0
	token.Wait()
	return token.Error()
}

// Closing the connection after x time
func (c *Client) Close() {
	c.client.Disconnect(disconnectTime)
}
