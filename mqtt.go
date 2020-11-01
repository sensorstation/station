package station

import (
	"fmt"
	"log"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var (
	mqttc MQTT.Client
)

func mqtt_connect() MQTT.Client {
	if config.Debug {
		MQTT.DEBUG = log.New(os.Stdout, "", 0)
		MQTT.ERROR = log.New(os.Stdout, "", 0)
	}

	server := config.Broker
	id := "sensorStation"
	connOpts := MQTT.NewClientOptions().AddBroker(server).SetClientID(id).SetCleanSession(true)

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return nil
	}
	return client
}
