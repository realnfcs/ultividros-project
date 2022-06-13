package savetemperedglass

type Output struct {
	Id     string `json:"id"`
	Status int    `json:"-"`
	Error  string `json:"error"`
}

func (*Output) Init(id string, status int, err error) *Output {
	if err != nil {
		return &Output{id, status, err.Error()}
	}

	return &Output{id, status, ""}
}
