package station

type Publisher interface {
	ID() string
	Publish() chan<- interface{}
}

type Src interface {
	Fetch() chan interface{}
}
