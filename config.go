package station

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
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

var (
	config Configuration
)

func init() {
	flag.StringVar(&config.Addr, "addr", "0.0.0.0:8011", "Address to listen for web connections")
	flag.StringVar(&config.App, "app", "../app/dist", "Directory for the web app distribution")
	flag.BoolVar(&config.Debug, "debug", false, "Start debugging")
	flag.BoolVar(&config.DebugMQTT, "debug-mqtt", false, "Debugging MQTT messages")
	flag.StringVar(&config.Filename, "config", "~/.config/sensors.json", "Where to read and store config")
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
func (c *Configuration) Save(fname string) error {

	jbuf, err := json.Marshal(c)
	if err != nil {
		log.Printf("[ERROR]: JSON marshaling config: %+v", err)
		return err
	}

	err = ioutil.WriteFile(fname, jbuf, 0644)
	if err != nil {
		log.Printf("[ERROR]: FILE writing config: %+v", err)
		return err
	}
	return err
}

// Load the file from the file corresponding to the fname parameter
func (c *Configuration) Load(fname string) error {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Printf("[ERROR]: failed to read file %s, %v", fname, err)
		return err
	}

	err = json.Unmarshal(buf, c)
	if err != nil {
		log.Printf("[ERROR]: failed to read file %s, %v", fname, err)
		return err
	}

	return err
}
