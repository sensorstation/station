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
		ID:   "0xdeadcafe", // MUST get MAC Address
		Addr: cfg.Addr,
		Done: make(chan string),
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
	log.Println("Starting reader / publishers ", s.Addr)
	for _, p := range s.Publishers {
		p.Publish(s.Done)
	}

	log.Println("Starting station Web and REST server on ", s.Addr)
	return http.ListenAndServe(s.Addr, nil)
}
