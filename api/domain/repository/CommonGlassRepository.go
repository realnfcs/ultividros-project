package repository

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// A interface CommonGlassRepository provém os métodos necessários para
// que o repositório a ser usado realiza o objetivo do usecase
type CommonGlassRepository interface {
	GetCommonGlass(string) (*entities.CommonGlass, int, error)
	SaveCommonGlass(entities.CommonGlass) (string, int, error)
}
