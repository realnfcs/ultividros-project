// Pacote responsável pela o usecase GetUser que executa
// a ação de pegar todas os usuários e retornar os dados ao
// cliente
package getusers

import "github.com/realnfcs/ultividros-project/api/domain/repository"

// Usecase responsável pela obtenção de todos os usuários de um
// repositório voltando um ponteiro de array de usuários
type GetUsers struct {
	UserRepository repository.UserRepository
}

func (g *GetUsers) Execute() *Output {
	e, status, err := g.UserRepository.GetUsers()
	return new(Output).Init(e, status, err)
}
