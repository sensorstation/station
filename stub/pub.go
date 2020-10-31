package main

import (
	"log"
	"time"
)

// Publisher represents a single stream of data (TimeSeries) emitted
// from a station.
type Publisher struct {
	Path   string
	Period time.Duration
}

// Publish will start producing data from the given data producer via
// the q channel returned to the caller. The caller lets Publish know
// to stop sending data when it receives a communication from the done channel
func (p *Publisher) Publish(done chan string) (q chan Data) {
	q = make(chan Data)
	ticker := time.NewTicker(sm.Period)

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-done:
				sm.Stop()
				log.Println("Random Data recieved a DONE, returning")
				break

			case <-ticker.C:
				f := sm.getData()
				d := DataValue{f}
				q <- d
			}
		}
	}()

	return q
}

func (sm *SoilMoisture) Stop() {

}
