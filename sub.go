package station

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MsgHandler interface {
	MsgHandler(c mqtt.Client, m mqtt.Message)
}

// Sub recieves messages on a certain channel and performs the actions
// according to those messages.
type Subscriber struct {
	Path string
	MsgHandler
}
