package main

import (
	"io"
	"log"
	"time"
)

// Publisher represents a single stream of data (TimeSeries) emitted
// from a station.
type Publisher struct {
	Path   string
	Period time.Duration
	Subs   []chan interface{}
	io.Reader
}

func NewPublisher(p string, r io.Reader) (pub *Publisher) {
	pub = &Publisher{
		Path:   p,
		Period: 5 * time.Second,
	}
	return pub
}

// Publish will start producing data from the given data producer via
// the q channel returned to the caller. The caller lets Publish know
// to stop sending data when it receives a communication from the done channel
func (p *Publisher) Publish(done chan string) {
	ticker := time.NewTicker(sm.Period)

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-done:
				p.Stop()
				log.Println("Random Data recieved a DONE, returning")
				break

			case <-ticker.C:
				var buf []byte
				n, err := p.Read(buf)
				if err != nil {
					log.Println("buf: " + buf)
					mqttcli.Publish(p.Path, byte(0), false, buf)
				}
			}
		}
	}()
}
