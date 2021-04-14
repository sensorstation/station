package main

import (
	"flag"
	"time"
)

// Globals
var (
	config Configuration
	mesh MeshNetwork
	tstempc Timeseries
	tshumid Timeseries
)

func init() {
	flag.StringVar(&config.Addr, "addr", "0.0.0.0:8011", "Address to listen for web connections")
	flag.StringVar(&config.App, "app", "../app/dist", "Directory for the web app distribution")
	flag.StringVar(&config.Broker, "broker", "tcp://localhost:1883", "Address of MQTT broker")
	flag.BoolVar(&config.Debug, "debug", false, "Start debugging")
	flag.BoolVar(&config.DebugMQTT, "debug-mqtt", false, "Debugging MQTT messages")
	flag.BoolVar(&config.FakeWS, "fake-ws", false, "Fake websocket data")
	flag.StringVar(&config.Filename, "config", "~/.config/sensors.json", "Where to read and store config")
	flag.BoolVar(&config.GPIO, "gpio", false, "Ignore the GPIO")

	mesh = MeshNetwork{
		Nodes: make(map[string]*MeshNode),
	}
	tstempc = Timeseries{
		Vals: make(map[time.Time]float64, 256),
	}
	tshumid = Timeseries{
		Vals: make(map[time.Time]float64, 256),
	}
}

func main() {

	// Parse command line argumens and update the config as appropriate
	flag.Parse()

	// Create the state configuration for this station.
	cfg := GetConfig()

	// Now create the station based on the given configuration
	st := NewStation(&cfg)

	// Register our REST callbacks, specifically answer to pings
	st.Register("/ws", WServ)
	st.Register("/ping", Ping{})
	st.Register("/config", Configuration{})

	// Subscribe to MQTT channels
	st.Subscribe("mesh/+/toCloud", ToCloudCB)

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
