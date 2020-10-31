/*

 */
package station

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Station struct {
	ID   string // MAC address
	Addr string

	*http.Server
}

func NewStation(cfg Config) (s *Station) {
	s = &Station{Addr: cfg.Addr}

	if err := mainImpl(); err != nil {
		fmt.Fprintf(os.Stderr, "gpio-list: %s.\n", err)
		os.Exit(1)
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
