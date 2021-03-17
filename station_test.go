package main

import "testing"

func TestStation(t *testing.T) {
	st := NewStation(&config)
	if st.Addr != "0.0.0.0:8011" {
		t.Errorf("station.Addr incorrect Expected (0.0.0.0:8011) got (%s)", st.Addr)
	}
}
