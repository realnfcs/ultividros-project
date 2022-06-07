package usecases

import (
	"github.com/NicolasSales0101/ultividros-project/api/domain/entities"
	"github.com/NicolasSales0101/ultividros-project/api/domain/repository"
)

// Usecase responsável pela salvamento do vidro temperado em um
// repositório passado por meio da inversão de dados
type SaveTemperedGlass struct {
	TemperedGlassRepository repository.TemperedGlassRepository
}

// Método que executa o procedimento de salvamento do vidro temperado
func (s *SaveTemperedGlass) Execute(e entities.TemperedGlass) error {
	return s.TemperedGlassRepository.SaveTemperedGlass(e)
}
