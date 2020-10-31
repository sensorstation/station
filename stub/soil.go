package main

import (
	"log"
	"time"
)

// SoilMoisture will collected a new piece of moisture data
// every period and transmit it to the given mqtt channel
type SoilMoisture struct {
	Period time.Duration // in milli-seconds
	Path   string        // mqtt & web endpoints to publish this data
}

// NewPeridocRandomDataGenerator will create a random number generator that
// publishes a new number every period
func NewSoilMoisture(p time.Duration, path string) *SoilMoisture {
	return &SoilMoisture{p, path}
}

// Publish will start producing data from the given data producer via
// the q channel returned to the caller. The caller lets Publish know
// to stop sending data when it receives a communication from the done channel
func (sm *SoilMoisture) Publish(done chan string) (q chan Data) {
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

func (sm *SoilMoisture) getData() float64 {
	return 55.5
}
