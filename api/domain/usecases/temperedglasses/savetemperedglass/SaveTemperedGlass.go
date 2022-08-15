// Pacote responsável pela o usecase SaveTemperedGlass que executa
// a ação de salvamento de um vidro temperado e retorna os dado de
// resposta ao cliente
package savetemperedglass

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável pela salvamento do vidro temperado em um
// repositório passado por meio da inversão de dados
type SaveTemperedGlass struct {
	TemperedGlassRepository repository.TemperedGlassRepository
	UserRepository          repository.UserRepository
}

// Método que executa o procedimento de salvamento do vidro temperado
func (s *SaveTemperedGlass) Execute(i Input) *Output {

	ocup, err := s.UserRepository.VerifyOccupation(i.UserId)
	if err != nil {
		return new(Output).Init("", 400, err)
	}

	if !ocup {
		return new(Output).Init("", 401, errors.New("unauthorized"))
	}

	id, status, err := s.TemperedGlassRepository.SaveTemperedGlass(*i.ConvertToTempGlss())
	return new(Output).Init(id, status, err)
}
