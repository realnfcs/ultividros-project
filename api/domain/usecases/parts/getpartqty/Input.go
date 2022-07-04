package getpartqty

// Usecase Input Port responsável pelos dados que entrarão
type Input struct {
	Id string
}

func (*Input) Init(id string) *Input {
	return &Input{id}
}
