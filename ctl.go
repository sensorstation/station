package station

type Controller struct {
	Setter
}

func NewController(p string, d Setter) *Controller {
	c := &Controller{
		Setter: d,
	}
	return c
}
