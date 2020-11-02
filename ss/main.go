package main

import (
	"flag"

	station ".."
)

// Globals
var (
	config station.Configuration
)

func main() {

	// Parse command line argumens and update the config as appropriate
	flag.Parse()

	// Create the state configuration for this station.
	cfg := station.GetConfig()

	// Now create the station based on the given configuration
	st := station.NewStation(&cfg)

	// Register our REST callbacks, specifically answer to pings
	st.Register("/ping", station.Ping{})
	st.Register("/config", station.Configuration{})

	// Register our publishers with their respective readers
	st.AddPublisher("data/cafedead/rando", station.NewRando())
	st.AddPublisher("data/cafedead/humidity", station.NewRando())
	st.AddPublisher("data/cafedead/solar", station.NewRando())
	st.AddPublisher("data/cafedead/soil", station.NewRando())
	st.AddPublisher("data/cafedead/tempf", station.NewRando())

	st.Subscribe("ctl/cafedead/pump", station.GetRelay("GPIO3"))
	st.Subscribe("ctl/cafedead/light", station.GetRelay("GPIO4"))
	st.Subscribe("ctl/cafedead/heater", station.GetRelay("GPIO8"))
	st.Subscribe("ctl/cafedead/fan", station.GetRelay("GPIO11"))

	st.Start()
}
