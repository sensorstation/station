package station

import (
	"net/http"
)

// WebServerIface allows apps to register callbacks to HTTP queries,
// including HTML5, REST and Websockets.
type Web interface {
	Register(path string, callback http.Handler)
	ServeFile(path string)
	Start(addr string)
	Stop()
}

// WWW servers up HTML5, REST
type WWW struct {
	*http.Server
}

func (w *WWW) Register(p string, h http.Handler) {
	http.Handle(p, h)
}

func (www *WWW) Start(addr string) error {
	ping := Ping{}
	www.Register("/ping", ping)
	return http.ListenAndServe(addr, nil)
}
