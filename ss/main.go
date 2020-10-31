package main

import (
	"flag"

	station ".."
)

// Globals
var (
	config station.Config
)

func init() {
	flag.StringVar(&config.Addr, "addr", ":1234", "Address to listen for HTTP connections")
	flag.StringVar(&config.Broker, "broker", "10.24.10.10", "Address of the message broker")
	flag.BoolVar(&config.IgnoreGPIO, "ignore-gpio", false, "Ignore the GPIO on this host")
}

func main() {

	// Parse command line argumens and update the config as appropriate
	flag.Parse()

	// Create the state configuration for this station.
	cfg := station.Config{Addr: ":1234"}

	// Now create the station based on the given configuration
	st := station.NewStation(cfg)

	// Register our REST callbacks, specifically answer to pings
	st.Register("/ping", station.Ping{})

	// Now start our station server
	st.Start()
}
