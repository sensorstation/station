package main

import (
	"flag"
	"fmt"
	"log"
	"time"
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

	go web()
	listen_loop()
}

func listen_loop() {

	// TODO: create table of publishers
	done := make(chan string)
	rando := NewPeridocRandomDataGenerator(1*time.Second, "rand")
	randoQ := rando.Publish(done)

	rand2 := NewPeridocRandomDataGenerator(1*time.Second, "rand2")
	rand2Q := rand2.Publish(done)

	soil := NewSoilMoisture(1*time.Second, "soil")
	soilQ := soil.Publish(done)

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
