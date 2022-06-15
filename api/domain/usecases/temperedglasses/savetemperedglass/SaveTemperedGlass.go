package savetemperedglass

import (
	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável pela salvamento do vidro temperado em um
// repositório passado por meio da inversão de dados
type SaveTemperedGlass struct {
	TemperedGlassRepository repository.TemperedGlassRepository
}

// Método que executa o procedimento de salvamento do vidro temperado
func (s *SaveTemperedGlass) Execute(i Input) *Output {
	id, status, err := s.TemperedGlassRepository.SaveTemperedGlass(*i.ConvertToTempGlss())
	return new(Output).Init(id, status, err)
}
