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
	st.Register("/ws", station.WServ)
	st.Register("/ping", station.Ping{})
	st.Register("/config", station.Configuration{})

	st.Subscribe("mesh/+/toCloud", station.ToCloudCB)

	// ----------------------------------------------------------
	// Register our publishers with their respective readers
	// ----------------------------------------------------------
	// st.AddPublisher("data/cafedead/soil", station.NewRando())

	// ----------------------------------------------------------
	// Register our local controls
	// ----------------------------------------------------------
	// rel1 := station.GetRelay("GPIO16")
	// st.Subscribe("ctl/cafedead/pump", rel1.MessageHandler)

	// ----------------------------------------------------------
	// Register the apps
	// ----------------------------------------------------------
	// t0 := station.GetToggle("ctl/cafedead/pump")
	// t0.Subscribe("data/cafedead/soil", t0.MessageHandler)
	//  st.AddApplication(t0)
	st.Start()
}
