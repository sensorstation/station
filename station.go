/*

 */
package station

import (
	"log"
	"net/http"
)

type Station struct {
	ID   string // MAC address
	Addr string

	*http.Server
}

// NewStation creates a station that did not previously exist.
// The ID will be populated with the MAC address of this node
func NewStation(cfg *Configuration) (s *Station) {
	s = &Station{
		ID:   "0xdeadcafe",
		Addr: cfg.Addr,
	}
	return s
}

func (s *Station) Register(p string, h http.Handler) {
	http.Handle(p, h)
}

func (s *Station) Start() error {
	log.Println("Starting Web server on ", s.Addr)
	return http.ListenAndServe(s.Addr, nil)
}
