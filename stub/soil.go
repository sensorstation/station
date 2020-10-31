package main

import (
	"fmt"
	"time"
)

// SoilMoisture will collected a new piece of moisture data
// every period and transmit it to the given mqtt channel
type SoilMoisture struct {
	Publisher
}

// NewPeridocRandomDataGenerator will create a random number generator that
// publishes a new number every period
func NewSoilMoisture() *SoilMoisture {
	return &SoilMoisture{
		Path:   "data/cafedead/soil",
		Period: time.Second * 5,
	}
}

func (sm *SoilMoisture) Read(b []byte) (n int, err error) {

	return 0.0, fmt.Errorf("TODO SoilMoisture READ")
}
