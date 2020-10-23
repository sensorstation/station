package station

import (
	"fmt"
	"log"

	"periph.io/x/periph"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

type GPIO struct {
	*periph.State
}

func (g *GPIO) Init() bool {
	var err error

	g.State, err = periph.Init()
	if err != nil {
		log.Fatalf("failed to init peripf: %v", err)
		return false
	}

	// it seems we were not able to load GPIO, hence we can't do
	// some stuff...
	if len(g.State.Loaded) == 0 {
		return false
	}
	return true
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
}
