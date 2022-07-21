package repository

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// A interface CommonGlassRepository provém os métodos necessários para
// que o repositório a ser usado realiza o objetivo do usecase
type CommonGlassRepository interface {
	GetArea(string) (map[string]float32, error)

	IncreaseArea(string, float32, float32) error
	ReduceArea(string, float32, float32) error

	GetCommonGlass(string) (*entities.CommonGlass, int, error)
	GetCommonGlasses() (*[]entities.CommonGlass, int, error)
	SaveCommonGlass(entities.CommonGlass) (string, int, error)
	PatchCommonGlass(entities.CommonGlass) (string, int, error)
	DeleteCommonGlass(entities.CommonGlass) (int, error)
}
