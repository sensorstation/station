package sensors

// Subscriber listens for data on the given channel, collects the data
// and does whatever it pleases with the data
type Subscriber interface {
	Channel() string
	Subscribe() <-chan interface{}
}
