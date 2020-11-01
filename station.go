//
// station is a program
//
package station

import (
	"log"
	"net/http"
)

type Station struct {
	ID   string // MAC address
	Addr string
	*http.Server
	Publishers map[string]*Publisher

	Done chan string
}

// NewStation creates a station that did not previously exist.
// The ID will be populated with the MAC address of this node
func NewStation(cfg *Configuration) (s *Station) {
	s = &Station{
		ID:         "0xdeadcafe", // MUST get MAC Address
		Addr:       cfg.Addr,
		Publishers: make(map[string]*Publisher, 10),
		Done:       make(chan string),
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

	log.Println("Connect to our ")
	mqttc = mqtt_connect()
	if mqttc == nil {
		log.Fatal("Unable to connect to broker, TODO StandAlone mode")
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
func (s *Station) AddPublisher(path string, r DataReader) {
	s.Publishers[path] = NewPublisher(path, r)
}
