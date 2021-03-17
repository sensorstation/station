package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Sub recieves messages on a certain channel and performs the actions
// according to those messages.
type Subscriber struct {
	Path string
	mqtt.MessageHandler
}
