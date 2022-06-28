package repository

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// A interface UserRepository provém os métodos necessários para
// que o repositório a ser usado realiza o objetivo do usecase
type UserRepository interface {
	GetUsers() (*[]entities.User, int, error)
	SaveUser(entities.User) (string, int, error)
}
