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

func GetRelay(name string) (r Relay) {
	r = Relay{GetPinRW(name)}
	return r
}

func (rel *Relay) Set(l gpio.Level) {
	rel.Out(l)
}

// MessageHandler will accept messages from the mqtt Client to be processed
// accordingly
func (rel *Relay) MessageHandler(c mqtt.Client, m mqtt.Message) {
	t := m.Topic()
	p := m.Payload()
	str := string(p)

	log.Println("Topic: ", t, " payload: ", str)

	l := gpio.Low
	switch str {
	case "on", "1", "true", "t":
		l = gpio.High
	case "off", "0", "false", "f":
		l = gpio.Low
	default:
		log.Printf("ERROR: relay message handler unknown message: %+v", str)
		return
	}

	if p[0] == byte(1) {
		l = gpio.High
	}
	log.Printf("Setting relay %s to %q", rel.PinWR.Name(), l)
	rel.Set(l)
}

func (rel *Relay) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Fatal("TODO Implement the server HTTP protocol")
}
