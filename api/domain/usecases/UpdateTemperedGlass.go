package usecases

import (
	"github.com/NicolasSales0101/ultividros-project/api/domain/entities"
	"github.com/NicolasSales0101/ultividros-project/api/domain/repository"
)

// Usecase responsável pela atualização de um vidro temperado em um
// repositório passado por meio da inversão de dados
type UpdateTemperedGlass struct {
	TemperedGlassRepository repository.TemperedGlassRepository
}

func (u *UpdateTemperedGlass) Execute(e entities.TemperedGlass) error {
	return u.TemperedGlassRepository.UpdateTemperedGlass(e)
}
