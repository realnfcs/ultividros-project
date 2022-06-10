package repository

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// A interface TemperedGlassRepository provém os métodos necessários para
// que o repositório a ser usado realiza o objetivo do usecase
type TemperedGlassRepository interface {
	GetTemperedGlass(string) *entities.TemperedGlass
	GetTemperedGlasses() *[]entities.TemperedGlass
	SaveTemperedGlass(entities.TemperedGlass) (string, int, error)
	UpdateTemperedGlass(entities.TemperedGlass) (string, int, error)
	PatchTemperedGlass(entities.TemperedGlass) (string, int, error)
	DeleteTemperedGlass(entities.TemperedGlass) (int, error)
}
