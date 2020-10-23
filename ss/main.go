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

	flag.Parse()

	st := station.Station{Addr: config.Addr}
	st.Register("/ping", station.Ping{})
	st.Start()
}
