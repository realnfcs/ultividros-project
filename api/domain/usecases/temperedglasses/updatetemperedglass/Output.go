package updatetemperedglass

type Output struct {
	Id     string
	Status int
	Error  error
}

func (*Output) Init(id string, status int, err error) *Output {
	return &Output{id, status, err}
}
