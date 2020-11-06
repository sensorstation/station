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

	if config.GPIO {
		
		// ----------------------------------------------------------
		// Register our publishers with their respective readers
		// ----------------------------------------------------------
		st.AddPublisher("data/cafedead/rando", station.NewRando())
		st.AddPublisher("data/cafedead/humidity", station.NewRando())
		st.AddPublisher("data/cafedead/solar", station.NewRando())
		st.AddPublisher("data/cafedead/soil", station.NewRando())
		st.AddPublisher("data/cafedead/tempf", station.NewRando())

		// ----------------------------------------------------------
		// Register our local controls
		// ----------------------------------------------------------
		rel1 := station.GetRelay("GPIO16")
		st.Subscribe("ctl/cafedead/pump", rel1.MessageHandler)

		rel2 := station.GetRelay("GPIO7")
		st.Subscribe("ctl/cafedead/light", rel2.MessageHandler)

		rel3 := station.GetRelay("GPIO8")
		st.Subscribe("ctl/cafedead/heater", rel3.MessageHandler)

		rel4 := station.GetRelay("GPIO24")
		st.Subscribe("ctl/cafedead/fan", rel4.MessageHandler)
	}

	// ----------------------------------------------------------
	// Register the apps
	// ----------------------------------------------------------
	t0 := station.GetToggle("ctl/cafedead/pump")
	t0.Subscribe("data/cafedead/soil", t0.MessageHandler)
	st.AddApplication(t0)

	s := station.GetToggle("ctl/cafedead/light")
	s.Subscribe("data/cafedead/solar", s.MessageHandler)
	st.AddApplication(s)

	s2 := station.GetToggle("ctl/cafedead/heater")
	s2.Subscribe("data/cafedead/tempf", s2.MessageHandler)
	st.AddApplication(s2)

	s3 := station.GetToggle("ctl/cafedead/fan")
	s3.Subscribe("data/cafedead/humidity", s3.MessageHandler)
	st.AddApplication(s3)
	st.Start()
}
