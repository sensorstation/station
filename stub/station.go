package main

// Station represents a single addressable sensor station
// complete with an Inventory
type Station struct {
	MAC      []byte
	NetID    string
	Sensors  []Sensor
	Controls []Control
}

// Sensor is an input that can be periodic or it may be
// event driven.
type Sensor struct {
	Id      string
	Period  int
	Value   Data
	History []Data
	Channel string
}

// Control is an output that can possibly manipulate the outside
// world
type Control struct {
	Id      string
	Val     float32
	Channel string
}
