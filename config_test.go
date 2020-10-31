package station

import (
	"flag"
	"testing"
)

func TestConfig(t *testing.T) {
	flag.Parse()

	fname := "/tmp/sensor-test.json"

	// config will already been configured
	err := config.Save(fname)
	if err != nil {
		t.Error("saving config: ", err)
	}

	var c Configuration
	err = c.Load(fname)
	if err != nil {
		t.Error("reading config: ", err)
	}
}
