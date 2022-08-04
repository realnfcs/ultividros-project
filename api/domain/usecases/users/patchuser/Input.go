package patchuser

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// Usecase Input Port responsável pelos dados que entrarão
type Input struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Occupation string `json:"occupation"`
}

func (*Input) Init(e entities.User) *Input {
	return &Input{
		e.Id,
		e.Name,
		e.Email,
		e.Password,
		e.Occupation,
	}
}

// Método responsável em converter um input em uma entidade de peça
func (i *Input) ConvertToUser() *entities.User {
	return &entities.User{
		Id:         i.Id,
		Name:       i.Name,
		Email:      i.Email,
		Password:   i.Password,
		Occupation: i.Occupation,
	}
}
