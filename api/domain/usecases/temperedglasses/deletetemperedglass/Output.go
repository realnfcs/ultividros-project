package deletetemperedglass

// Usecase Output Port responsável pelos dados que serão retornados
type Output struct {
	status int
	err    error
}

func (*Output) Init(status int, err error) *Output {
	return &Output{status, err}
}
