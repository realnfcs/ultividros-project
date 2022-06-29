package getuser

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// Usecase Output Port responsável pelos dados que serão retornados
type Output struct {
	OutputData `json:"data"`
	Status     int    `json:"-"`
	Err        string `json:"error"`
}

// OutpurData respnsável pelo dado da entity em si que será passado pelas camadas externas
type OutputData struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Occupation string `json:"occupation"`
}

func (*Output) Init(e *entities.User, status int, err error) *Output {
	if e != nil {
		return &Output{
			OutputData{
				e.Id,
				e.Name,
				e.Email,
				e.Password,
				e.Occupation,
			},
			status,
			"",
		}
	}

	return &Output{OutputData{}, status, err.Error()}
}
