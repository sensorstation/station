package main

import "time"

type Timeseries struct {
	Vals map[time.Time]float64
}

func (ts *Timeseries) Add(v float64) {
	ts.Vals[time.Now()] = v
}

