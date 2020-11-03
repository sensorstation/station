package station

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Application interface {
	Name() string
	Subscribe(topic string, f mqtt.MessageHandler)
	// Register(path string)  for http
}

type App struct {
	name          string
	Subscriptions []string
}

func NewApp(name string) (app App) {
	app = App{
		name:          name,
		Subscriptions: make([]string, 5),
	}
	return app
}

func (app App) Name() string {
	return app.name
}

func (app App) Subscribe(topic string, f mqtt.MessageHandler) {
	qos := 0

	app.Subscriptions = append(app.Subscriptions, topic)
	if token := mqttc.Subscribe(topic, byte(qos), f); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		log.Printf("\tsubscribe token: %v", token)
	}
	log.Println("\t%s subscribed to %s ", app.Name, topic)
}
