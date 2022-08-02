// Pacote responsável pela o usecase SaveTemperedGlass que executa
// a ação de salvamento de um vidro comum e retorna os dado de
// resposta ao cliente
package savecommonglass

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável pela salvamento do vidro comum em um
// repositório passado por meio da inversão de dados
type SaveCommonGlass struct {
	CommonGlassRepository repository.CommonGlassRepository
	UserRepository        repository.UserRepository
}

// Método que executa o procedimento de salvamento do vidro comum
func (s *SaveCommonGlass) Execute(i Input) *Output {

	ocup, err := s.UserRepository.VerifyOccupation(i.UserId)
	if err != nil {
		return new(Output).Init("", 400, err)
	}

	if !ocup {
		return new(Output).Init("", 401, errors.New("unauthorized"))
	}

	id, status, err := s.CommonGlassRepository.SaveCommonGlass(*i.ConvertToComnGlss())
	return new(Output).Init(id, status, err)
}
