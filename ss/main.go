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

	rel1 := station.GetRelay("GPIO16")
	st.Subscribe("ctl/cafedead/pump", rel1.MessageHandler)

	rel2 := station.GetRelay("GPIO7")
	st.Subscribe("ctl/cafedead/light", rel2.MessageHandler)

	rel3 := station.GetRelay("GPIO8")
	st.Subscribe("ctl/cafedead/heater", rel3.MessageHandler)

	rel4 := station.GetRelay("GPIO24")
	st.Subscribe("ctl/cafedead/fan", rel4.MessageHandler)

	st.Start()
}
