package getpartqty

// Usecase Output Port responsável pelos dados que serão retornados
type Output struct {
	Id  string `json:"id"`
	Qty uint32 `json:"qty"`
	Err error  `json:"error"`
}

func (*Output) Init(id string, qty uint32, err error) *Output {
	return &Output{id, qty, err}
}
