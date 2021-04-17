//
// station is a program
//
package main

import (
	"log"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"periph.io/x/periph"
	"periph.io/x/periph/host"
)

type Station struct {
	ID   string // MAC address
	Addr string
	*http.Server
	*periph.State
	mqttc *mqtt.Client

	Publishers   map[string]*Publisher
	Subscribers  map[string]*Subscriber
	Applications map[string]Application

	Done chan string
}

// NewStation creates a station that did not previously exist.
// The ID will be populated with the MAC address of this node
func NewStation(cfg *Configuration) (s *Station) {
	s = &Station{
		ID:          "0xdeadcafe", // MUST get MAC Address - and WIFI SSID
		Addr:        cfg.Addr,
		Publishers:  make(map[string]*Publisher, 10),
		Subscribers: make(map[string]*Subscriber, 10),
		Applications: make(map[string]Application, 10),
		Done:        make(chan string),
	}

	mqtt_connect()
	if mqttc == nil {
		log.Fatalf("Failed to connect to MQTT broker")
	}

	var err error
	s.State, err = host.Init()
	if err != nil {
		log.Printf("Initializing GPIO failed - no GPIO")
		s.State = nil
	}
	return s
}

// Register to handle HTTP requests for particular paths in the
// URL or MQTT channel.
func (s *Station) Register(p string, h http.Handler) {
	http.Handle(p, h)
}

// Start the HTTP server and serve up the home web app and
// our REST API
func (s *Station) Start() error {

	log.Println("Connect to our MQTT broker: ", config.Broker)
	if mqttc == nil {
		log.Fatal("Unable to connect to broker, TODO StandAlone mode")
	}

	log.Println("Subscribers: ", len(s.Subscribers))
	for _, p := range s.Subscribers {
		log.Println("\t" + p.Path)
	}

	log.Println("Starting publishers: ", len(s.Publishers))
	for _, p := range s.Publishers {
		log.Println("\t" + p.Path)
		p.Publish(s.Done)
	}

	log.Println("Starting station Web and REST server on ", s.Addr)
	return http.ListenAndServe(s.Addr, nil)
}

// NewPublisher adds a publisher to the station, which will subsequently
// start publishing the data
func (s *Station) AddApplication(app Application) {
	name := app.Name()
	if s.Applications == nil {
		s.Applications = make(map[string]Application, 4)
	}
	s.Applications[name] = NewApp(name)
}

func (s *Station) Subscribe(path string, f mqtt.MessageHandler) {
	sub := &Subscriber{path, f}
	s.Subscribers[path] = sub

	qos := 0
	if token := mqttc.Subscribe(path, byte(qos), f); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		log.Printf("subscribe token: %v", token)
	}
	log.Println("Subscribed to ", path)
}
