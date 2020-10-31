package main

import (
	"encoding/json"
	"flag"
	"net/http"
)

type Configuration struct {
	Addr string
	App  string

	Debug     bool
	DebugMQTT bool
	DumpGPIO  bool
	Filename  string

	IgnoreGPIO  bool
	ShowSkipped bool
}

func init() {
	flag.StringVar(&config.Addr, "addr", "0.0.0.0:8011", "Address to listen for web connections")
	flag.StringVar(&config.App, "app", "../app/dist", "Directory for the web app distribution")

	flag.StringVar(&mesh.ID, "mesh-id", "mobrob", "MESH Network ID")
	flag.StringVar(&mesh.SSID, "ssid", "super-stuff", "SSID of the optional IP Router")
	flag.StringVar(&mesh.Pass, "pass", "wifi-pass", "Password of the wifi also used as mesh password")

	flag.BoolVar(&config.Debug, "debug", false, "Start debugging")
	flag.BoolVar(&config.Debug, "debug", false, "Start debugging")
	flag.BoolVar(&config.DebugMQTT, "debug-mqtt", false, "Debugging MQTT messages")
	flag.BoolVar(&config.DumpGPIO, "gpio", false, "Start debugging")
	flag.StringVar(&config.Filename, "config", "~/.config/sensors.json", "Where to read and store config")
	flag.BoolVar(&config.IgnoreGPIO, "ignore-gpio", false, "Ignore GPIO for computers without GPIO")
	flag.BoolVar(&config.ShowSkipped, "show-skipped", false, "Show skipped driver periphs")

}

// ServeHTTP allows a user to get and possibly set the configuration
func (c Configuration) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(config)

	case "POST", "PUT":
		// TODO
		http.Error(w, "Not Yet Supported", 401)
	}
}

// Save write the configuration to a file in JSON format
func (c *Configuration) Save(fname string) {

}

// Load the file from the file corresponding to the fname parameter
func (c *Configuration) Load(fname string) {
}
