package station

import (
	"io"
	"log"
	"time"

	"periph.io/x/periph/conn/gpio"
)

// Publisher periodically reads from an io.Reader then publishes that value
// to a corresponding channel
type Publisher struct {
	Path   string
	Period time.Duration
	Pin    gpio.PinIO
	io.Reader

	publishing bool
}

func NewPublisher(p string, r io.Reader) (pub *Publisher) {
	pub = &Publisher{
		Path:   p,
		Period: 5 * time.Second,
	}
	return pub
}

func (p *Publisher) Stop() {
	p.publishing = false
}

// Publish will start producing data from the given data producer via
// the q channel returned to the caller. The caller lets Publish know
// to stop sending data when it receives a communication from the done channel
func (p *Publisher) Publish(done chan string) {
	ticker := time.NewTicker(p.Period)

	go func() {
		defer ticker.Stop()
		p.publishing = true
		for p.publishing {
			select {
			case <-done:
				p.Stop()
				log.Println("Random Data recieved a DONE, returning")
				break

			case <-ticker.C:
				var buf []byte
				n, err := p.Read(buf)
				if err == nil {
					log.Printf("ERROR publish [%s] value %+v", p.Path, buf)
					continue
				}
				if config.Debug {
					log.Printf("chan %s published %d bytes: %+v", p.Path, n, buf)
				}
				mqttc.Publish(p.Path, byte(0), false, buf)
			}
		}
	}()
}
