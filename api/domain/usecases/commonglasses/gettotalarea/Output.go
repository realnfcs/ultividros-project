package gettotalarea

// Usecase Output Port responsável pelos dados que serão retornados
type Output struct {
	Id              string  `json:"id"`
	WidthAvailable  float32 `json:"width_available"`
	HeightAvailable float32 `json:"height_available"`
	Err             error   `json:"error"`
}

func (*Output) Init(id string, widthAvailable, heightAvailable float32, err error) *Output {
	return &Output{id, widthAvailable, heightAvailable, err}
}
