package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

// --------------------------------------------------------------------
// PeriodicRandomDataGenerator - periodically generates random data
// --------------------------------------------------------------------

// PeriodicRandomData will collected a new random piece of data
// every period and transmit it to the given mqtt channel
type PeriodicRandomDataGenerator struct {
	Period time.Duration // in milli-seconds
	Path   string        // mqtt & web endpoints to publish this data
}

// NewPeridocRandomDataGenerator will create a random number generator that
// publishes a new number every period
func NewPeridocRandomDataGenerator(p time.Duration, path string) *PeriodicRandomDataGenerator {
	return &PeriodicRandomDataGenerator{p, path}
}

// Publish will start producing data from the given data producer via
// the q channel returned to the caller. The caller lets Publish know
// to stop sending data when it receives a communication from the done channel
func (rd *PeriodicRandomDataGenerator) Publish(done chan string) (q chan Data) {
	q = make(chan Data)
	ticker := time.NewTicker(rd.Period)

	rand.Seed(int64(os.Getpid()))
	go func() {
		for {
			select {
			case <-done:
				log.Println("Random Data recieved a DONE, returning")
				return

			case <-ticker.C:
				f := rand.Float64()
				d := DataValue{f}
				q <- d
			}
		}
	}()

	return q
}
