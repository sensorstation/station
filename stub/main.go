package main

import (
	"flag"
	"fmt"
	"log"
)

const (
	PIN_SOIL  = 32
	PIN_HUM   = 12
	PIN_SOLAR = 11
	PIN_TEMPF = 36
)

var (
	config Configuration // config.go
	ws     WS
)

func main() {

	log.Println("Parse the flags")
	flag.Parse()

	log.Println("Read the config file")
	config.Load("~/.config/station.json")

	log.Println("Startup the GPIOs")
	if !config.IgnoreGPIO {
		g := GetPins()
		if g == nil {
			log.Println("GPIO not found")
		}
	} else {
		log.Println("Ignoring GPIO")
	}

	// Get the readers
	log.Println("Create all of our Subscribers")

	go web()
	listen_loop()
}

func listen_loop() {

	var d Data
	mqttcli := mqtt()
	for {
		select {
		case d = <-randoQ:
			f := d.String()
			mqttcli.Publish(rando.Path, byte(0), false, f)

		case d = <-rand2Q:
			f := d.String()
			mqttcli.Publish(rand2.Path, byte(0), false, f)

		case d = <-soilQ:
			f := d.String()
			mqttcli.Publish(soil.Path, byte(0), false, f)

		case <-done:
			log.Println("MQTT Recieved a done, returning")
			return
		}

		if config.DebugMQTT {
			fmt.Printf("mqtt %s -> %6.2f\n", soil.Path, d.Float64())
		}
	}
}
