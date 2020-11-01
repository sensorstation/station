package station

import (
	"log"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

var readers map[string]*PinReader

type PinReader struct {
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

func (p PinReader) Read(buf []byte) (n int, err error) {
	l := p.PinIO.Read()
	if l == gpio.Low {
		buf = append(buf, 0)
	} else {
		buf = append(buf, 1)
	}
	return len(buf), nil
}
