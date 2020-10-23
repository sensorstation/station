/*

 */
package station

import "net/http"

type Station struct {
	ID   string // MAC address
	Addr string

	*http.Server
}

func (s *Station) Register(p string, h http.Handler) {
	http.Handle(p, h)
}

func (s *Station) Start() error {
	return http.ListenAndServe(s.Addr, nil)
}
