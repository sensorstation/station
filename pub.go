package station

// Publisher advertises data over our Pub/Sub protocol (most likely
// MQTT). The publisher has to retrieve the data from somewhere.
type Publisher interface {
	Channel() string
	Publish() chan<- interface{}
}
