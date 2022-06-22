package repository

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// A interface PartRepository provém os métodos necessários para
// que o repositório a ser usado realiza o objetivo do usecase
type PartRepository interface {
	GetParts() (*[]entities.Part, int, error)
	SavePart(entities.Part) (string, int, error)
}
