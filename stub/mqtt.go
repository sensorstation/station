package main

import (
	"fmt"
	"log"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func mqtt() MQTT.Client {
	if config.Debug {
		MQTT.DEBUG = log.New(os.Stdout, "", 0)
		MQTT.ERROR = log.New(os.Stdout, "", 0)
	}

	server := "tcp://10.24.10.10:1883" // XXX move to config
	id := "sensorStation"
	connOpts := MQTT.NewClientOptions().AddBroker(server).SetClientID(id).SetCleanSession(true)

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return nil
	}
	mesh_subscribe(client)
	return client
}
