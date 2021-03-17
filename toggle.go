package main

import (
	"log"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Toggle struct {
	Control string
	App
}

func GetToggle(ctl string) (tog *Toggle) {
	tog = &Toggle{Control: ctl}
	tog.App = NewApp("Sprinkler")
	return tog
}

func (tog *Toggle) MessageHandler(c mqtt.Client, m mqtt.Message) {
	top := m.Topic()
	pay := m.Payload()
	str := string(pay)

	log.Println("Topic: ", top, " payload: ", str)

	var err error
	var f float64
	if f, err = strconv.ParseFloat(str, 64); err != nil {
		log.Printf("Error converting %s to float: %v", str, err)
		return
	}
	log.Printf("App MQTT [IN] topic: %s - %v", top, f)

	cmd := "off"
	if f > 0.50 {
		cmd = "on"
	}
	if tok := mqttc.Publish(tog.Control, byte(0), false, cmd); tok == nil {
		log.Printf("I have a NULL token: %s, %+v", top, cmd)
	}
	log.Printf("Spinklers told %s to turn %+v", tog.Control, cmd)
}
