package main

import (
	"math/rand"
)

// --------------------------------------------------------------------
// PeriodicRandomDataGenerator - periodically generates random data
// --------------------------------------------------------------------

// PeriodicRandomData will collected a new random piece of data
// every period and transmit it to the given mqtt channel
type Rando struct {
	Publisher
}

func NewRando(path string) (r *Rando) {
	r = &Rando{}
	r.Publisher = NewPublisher(path, r)
	return r
}

func (p *Rando) Read(buf []byte) (n int, err error) {
	f := rand.Float64()
	return []byte(f)
}
