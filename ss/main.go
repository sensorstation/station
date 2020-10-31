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
	cfg := station.Configuration{Addr: ":1234"}

	// Now create the station based on the given configuration
	st := station.NewStation(&cfg)

	// Register our REST callbacks, specifically answer to pings
	st.Register("/ping", station.Ping{})
	st.Register("/config", station.Configuration{})

	// Now start our station server
	st.Start()
}
