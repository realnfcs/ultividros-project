package saveuser

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// Usecase Input Port responsável pelos dados que entrarão
type Input struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Occupation string `json:"occupation"`
}

func (*Input) Init(e entities.User) *Input {
	return &Input{
		e.Name,
		e.Email,
		e.Password,
		e.Occupation,
	}
}

// Método que converte um input na entidade User
func (i *Input) ConvertToUser() *entities.User {
	return &entities.User{
		"",
		i.Name,
		i.Email,
		i.Password,
		i.Occupation,
	}
}
