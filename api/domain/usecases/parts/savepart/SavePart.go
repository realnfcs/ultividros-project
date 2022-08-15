// Pacote responsável pela o usecase SavePart que executa
// a ação de salvamento de uma peça e retorna os dado de
// resposta ao cliente
package savepart

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável pela salvamento de uma peça em um
// repositório passado por meio da inversão de dados
type SavePart struct {
	PartRepository repository.PartRepository
	UserRepository repository.UserRepository
}

// Método que executa o procedimento de salvamento da peça
func (s *SavePart) Execute(i Input) *Output {

	ocup, err := s.UserRepository.VerifyOccupation(i.UserId)
	if err != nil {
		return new(Output).Init("", 400, err)
	}

	if !ocup {
		return new(Output).Init("", 401, errors.New("unauthorized"))
	}

	id, status, err := s.PartRepository.SavePart(*i.ConvertToPart())
	return new(Output).Init(id, status, err)
}
