package main

import (
	"strconv"
	"time"
)

// --------------------------------------------------------------------
// Data Interfaces
// --------------------------------------------------------------------

// Data interface
type Data interface {
	Float64() float64
	String() string
}

// DataGenerator
type DataGenarator interface {
	Publish(cmd chan string) chan Data
}

// --------------------------------------------------------------------
// DataValue - Simply a float64 value
// --------------------------------------------------------------------

// DataVale is just a float64 that satisfies the Data interface
type DataValue struct {
	Val float64
}

// Float64 returns the value of this data
func (d DataValue) Float64() float64 {
	return d.Val
}

// Float64 returns the value of this data
func (d DataValue) String() string {
	s64 := strconv.FormatFloat(d.Val, 'f', 2, 64)
	return s64
}

// --------------------------------------------------------------------
// DataTS - DataValue and TimeStamp
// --------------------------------------------------------------------

// DataTS is composed of a DataValue and TimeStamp
type DataTS struct {
	DataValue
	time.Time
}

// Float64 returns the value of this value
func (d DataTS) Float64() float64 {
	return d.Val
}

// String returns the string of this value
func (d DataTS) String() string {
	return d.DataValue.String()
}

// Float64 returns the value of this value
func (d DataTS) TimeStamp() time.Time {
	return d.Time
}
