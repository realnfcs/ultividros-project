// Pacote responsável pela o usecase SaveUser que executa
// a ação de salvamento de um usuário e retorna os dado de
// resposta ao cliente
package saveuser

import "github.com/realnfcs/ultividros-project/api/domain/repository"

// Usecase responsável pela salvamento de um usuário em um
// repositório passado por meio da inversão de dados
type SaveUser struct {
	UserRepository repository.UserRepository
}

func (s *SaveUser) Execute(i Input) *Output {
	id, status, err := s.UserRepository.SaveUser(*i.ConvertToUser())
	return new(Output).Init(id, status, err)
}
