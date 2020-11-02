package station

type Getter interface {
	Get() interface{}
}

type Setter interface {
	Set(d interface{})
}
