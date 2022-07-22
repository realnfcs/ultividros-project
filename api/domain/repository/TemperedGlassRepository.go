package repository

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// A interface TemperedGlassRepository provém os métodos necessários para
// que o repositório a ser usado realiza o objetivo do usecase
type TemperedGlassRepository interface {
	GetTempGlssQty(string) (uint32, error)

	IncreaseQuantity(string, uint32) error
	ReduceQuantity(string, uint32) error

	GetTemperedGlass(string) (*entities.TemperedGlass, int, error)
	GetTemperedGlasses() (*[]entities.TemperedGlass, int, error)
	SaveTemperedGlass(entities.TemperedGlass) (string, int, error)
	UpdateTemperedGlass(entities.TemperedGlass) (string, int, error)
	PatchTemperedGlass(entities.TemperedGlass) (string, int, error)
	DeleteTemperedGlass(entities.TemperedGlass) (int, error)
}
