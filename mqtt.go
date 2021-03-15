package station

import (
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

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

func ToCloudCB(mc mqtt.Client, msg mqtt.Message) {
	fmt.Printf("MC: %+v - MSG: %+v\n", mc, msg);
}
