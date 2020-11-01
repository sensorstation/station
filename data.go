package station

type DataReader interface {
	FetchData() interface{}
}

type DataSetter interface {
	SetData(d interface{}) error
}

