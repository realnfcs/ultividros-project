// Pacote responsável pela o usecase GetUser que executa
// a ação de pegar um usuário e retornar os dados de acordo
// com o ID passado
package getuser

import "github.com/realnfcs/ultividros-project/api/domain/repository"

// Usecase responsável pela obtenção de um único usuário de um
// repositório voltando um ponteiro da entidade com o mesmo ID passado
// no parâmetro
type GetUser struct {
	UserRepository repository.UserRepository
}

func (g *GetUser) Execute(i Input) *Output {
	e, status, err := g.UserRepository.GetUser(i.ID)
	return new(Output).Init(e, status, err)
}
