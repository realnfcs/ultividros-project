package deletepart

// Usecase Output Port responsável pelos dados que serão retornados
type Output struct {
	Status int    `json:"-"`
	Err    string `json:"error"`
}

func (*Output) Init(status int, err error) *Output {
	if err != nil {
		return &Output{status, err.Error()}
	}

	return &Output{status, ""}
}
