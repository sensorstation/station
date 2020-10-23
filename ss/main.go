package main

import (
	"flag"

	sensors ".."
)

// Globals
var (
	config sensors.Config
)

func init() {
	flag.StringVar(&config.Addr, "addr", ":1234", "Address to listen for HTTP connections")
	flag.StringVar(&config.Broker, "broker", "10.24.10.10", "Address of the message broker")
	flag.BoolVar(&config.IgnoreGPIO, "ignore-gpio", false, "Ignore the GPIO on this host")
}

func main() {

	flag.Parse()

	sensors.Init()

	// Start web server
	www := sensors.WWW{}
	www.Register("/ping", ping{})
	www.Start("0.0.0.0:1234")
}