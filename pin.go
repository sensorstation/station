package station

import (
	"log"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

var readers map[string]*PinReader

type PinWR struct {
	gpio.PinIO
}

func GetPinReader(name string) (p PinReader) {
	if p.PinIO = gpioreg.ByName(name); p.PinIO == nil {
		log.Fatalln("Could not find pin: ", name)
	}

	if err := p.In(gpio.Float, gpio.NoEdge); err != nil {
		log.Fatalln("Could not set input params on pin", name)
	}
	return p
}

func (p PinWR) FetchData() interface{} {
	var o bool
	l := p.PinIO.Read()
	if l == gpio.Low {
		o = false
	} else {
		o = true
	}
	return o
}

func (p PinWR) SetData(d interface{}) (err error) {

	return err
}

func (p PinWR) SetLevel(l gpio.Level) (err error) {

	
	return err
}
