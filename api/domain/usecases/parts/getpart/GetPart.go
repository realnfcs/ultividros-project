// Pacote responsável pela o usecase GetPart que executa
// a ação de pegar uma peça e retornar os dados de acordo
// com o ID passado
package getpart

import "github.com/realnfcs/ultividros-project/api/domain/repository"

type GetPart struct {
	PartRepository repository.PartRepository
}

// Usecase responsável pela obtenção de uma única peça de um
// repositório voltando um ponteiro da entidade com o mesmo ID passado
// no parâmetro
func (g *GetPart) Execute(i Input) *Output {
	e, status, err := g.PartRepository.GetPart(i.ID)
	return new(Output).Init(e, status, err)
}
