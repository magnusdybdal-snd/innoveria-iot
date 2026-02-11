package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type consumer struct {
	client mqtt.Client
}


