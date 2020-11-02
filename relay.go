package station

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"periph.io/x/periph/conn/gpio"
)

type Relay struct {
	PinWR
}

func GetRelay(name string) (p PinWR) {
	p = GetPinRW(name)
	return p
}

func (r *Relay) Set(l gpio.Level) {
	r.Set(l)
}

func (r *Relay) MessageHandler(c mqtt.Client, m mqtt.Message) {

}
