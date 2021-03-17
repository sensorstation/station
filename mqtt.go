package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ToCloudMsg struct {
	Addr	string `json:"addr"`
	Type	string `json:"type"`
	Data	map[string]interface{}	`json:"data"`
}

var (
	mqttc mqtt.Client
)

func mqtt_connect() {
	if config.DebugMQTT {
		mqtt.DEBUG = log.New(os.Stdout, "", 0)
		mqtt.ERROR = log.New(os.Stdout, "", 0)
	}

	id := "sensorStation"
	connOpts := mqtt.NewClientOptions().AddBroker(config.Broker).SetClientID(id).SetCleanSession(true)
	mqttc = mqtt.NewClient(connOpts)
	if token := mqttc.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
}

// ToCloudCB is the callback when we recieve MQTT messages on the '/mesh/xxxxxx/toCloud' channel. 
func ToCloudCB(mc mqtt.Client, msg mqtt.Message) {
	if false {
		log.Printf("Incoming message topic: %s\n", msg.Topic());
	}

	paths := strings.Split(msg.Topic(), "/");
	if len(paths) != 3 {
		log.Fatal("Error unsupported path")
	}

	// Get the node ID
	rootid := paths[1]
	n := mesh.GetNode(rootid)
	if n == nil {
		log.Fatalln("GetNode returned nil for ", rootid)
	}

	var m ToCloudMsg
	err := json.Unmarshal(msg.Payload(), &m)
	if err != nil {
		log.Fatal("Failed to unmarshal payload")
	}

	// TODO create a fully configured node and schedule network topology updates.
	log.Printf("%2.0f: %20s -> %-20s Parent: %s\n", m.Data["layer"], m.Addr, m.Data["self"], m.Data["parent"])
}
