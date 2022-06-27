// Pacote responsável pela o usecase SavePart que executa
// a ação de salvamento de uma peça e retorna os dado de
// resposta ao cliente
package savepart

import "github.com/realnfcs/ultividros-project/api/domain/repository"

// Usecase responsável pela salvamento de uma peça em um
// repositório passado por meio da inversão de dados
type SavePart struct {
	PartRepository repository.PartRepository
}

// Método que executa o procedimento de salvamento da peça
func (s *SavePart) Execute(i Input) *Output {
	id, status, err := s.PartRepository.SavePart(*i.ConvertToPart())
	return new(Output).Init(id, status, err)
}
