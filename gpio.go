package sensors

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/periph"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host/rpi"
)

type GPIO struct {
	*periph.State
}

func (g *GPIO) Init() {
	var err error

	g.State, err = periph.Init()
	if err != nil {
		log.Fatalf("failed to init peripf: %v", err)
	}

	// it seems we were not able to load GPIO, hence we can't do
	// some stuff...
	if len(g.State.Loaded) == 0 {
		config.IgnoreGPIO = true
		return false
	}
}

func (g *GPIO) Dump() {
	fmt.Printf("Using drivers:\n")
	for _, driver := range g.State.Loaded {
		fmt.Printf("- %s\n", driver)
	}

	fmt.Printf("Drivers skipped:\n")
	for _, failure := range g.State.Skipped {
		fmt.Printf("- %s: %s\n", failure.D, failure.Err)
	}

	fmt.Printf("Drivers failed to load:\n")
	for _, failure := range g.State.Failed {
		fmt.Printf("- %s: %v\n", failure.D, failure.Err)
	}
	fmt.Print("GPIO pins available:\n")
	for _, p := range gpioreg.All() {
		fmt.Printf("- %7s: %4d - %s\n", p, p.Number(), p.Function())
	}
	return true
}

func blinker() {
	ticker := time.NewTicker(1000 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return

			case t := <-ticker.C:
				l := !gpio.Low
				rpi.P1_33.Out(l)
				if config.Debug {
					fmt.Println("Tick: ", t)
				}
			}
		}
	}()
}
