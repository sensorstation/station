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

	m := map[string]string{
		//{"data/cafedead/tempf": "GPIO5"},
		//{"data/cafedead/humidity": "GPIO6"},
		//{"data/cafedead/solar": "GPIO9"},
		//{"data/cafedead/soil": "GPIO10"},
	}

	for c, p := range m {
		st.Publishers[c] = station.NewPublisher(c, station.GetPinReader(p))
	}
	st.Start()
}
