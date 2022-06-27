package repository

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// A interface PartRepository provém os métodos necessários para
// que o repositório a ser usado realiza o objetivo do usecase
type PartRepository interface {
	GetPart(string) (*entities.Part, int, error)
	GetParts() (*[]entities.Part, int, error)
	SavePart(entities.Part) (string, int, error)
	PatchPart(entities.Part) (string, int, error)
	DeletePart(entities.Part) (int, error)
}
