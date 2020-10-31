package main

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/periph"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/rpi"
)

type Pins struct {
	*periph.State
}

var (
	pins *Pins
)

func GetPins() *Pins {
	if pins != nil {
		return pins
	}

	var err error

	pins = &Pins{}
	pins.State, err = host.Init()
	if err != nil {
		log.Printf("failed to initialize periph: %v", err)
		return nil
	}

	// it seems we were not able to load PINS, hence we can't do
	// some stuff...
	if len(pins.State.Loaded) == 0 {
		return nil
	}
	return pins
}

func (g *Pins) String() string {

	str := "Using drivers:\n"
	for _, driver := range g.State.Loaded {
		str += fmt.Sprintf("- %s\n", driver)
	}

	str += "Drivers skipped:\n"
	for _, failure := range g.State.Skipped {
		str += fmt.Sprintf("- %s: %s\n", failure.D, failure.Err)
	}

	str += "Drivers failed to load:\n"
	for _, failure := range g.State.Failed {
		str += fmt.Sprintf("- %s: %v\n", failure.D, failure.Err)
	}
	str += "GPIO pins available:\n"
	for _, p := range gpioreg.All() {
		str += fmt.Sprintf("- %7s: %4d - %s\n", p, p.Number(), p.Function())
	}
	return str
}

func blinker() {
	ticker := time.NewTicker(1000 * time.Millisecond)
	done := make(chan bool)
	l := gpio.Low

	go func() {
		for {
			select {
			case <-done:
				return

			case t := <-ticker.C:
				l := !l
				rpi.P1_33.Out(l)
				if config.Debug {
					fmt.Println("Tick: ", t)
				}
			}
		}
	}()
}
