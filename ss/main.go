package main

import (
	"flag"
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

	station.Init()

	// Start web server
	www := station.WWW{}
	www.Register("/ping", ping{})
	www.Start("0.0.0.0:1234")
}
