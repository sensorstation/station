package station

type Publisher interface {
	ID() string
	Publish() chan<- interface{}
}

type Src interface {
	Fetch() chan interface{}
}

// Publisher advertises data over our Pub/Sub protocol (most likely
// MQTT). The publisher has to retrieve the data from somewhere.
type Pub struct {
	id string
}

func (p *Pub) ID() string {
	return id
}
