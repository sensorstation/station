package station

import (
	"log"
	"net/http"

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

// MessageHandler will accept messages from the mqtt Client to be processed
// accordingly
func (r *Relay) MessageHandler(c mqtt.Client, m mqtt.Message) {
	t := m.Topic()
	p := m.Payload()

	log.Println("Topic: ", t, " payload: ", p)
	l := gpio.Low
	if p[0] == byte(1) {
		l = gpio.High
	}
	r.Set(l)
}

func (r *Relay) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
