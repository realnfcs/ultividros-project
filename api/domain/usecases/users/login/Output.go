package login

// Usecase Output Port responsável pelos dados que serão retornados
type Output struct {
	Token  string `json:"token"`
	Status int    `json:"-"`
	Err    string `json:"error"`
}

func (*Output) Init(token string, status int, err error) *Output {
	if err != nil {
		return &Output{
			Token:  token,
			Status: status,
			Err:    err.Error(),
		}
	}

	return &Output{
		Token:  token,
		Status: status,
		Err:    "",
	}
}
