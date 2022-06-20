// Pacote responsável pela o usecase GetCommonGlass que executa
// a ação de pegar um vidro comum e retornar os dados de acordo
// com o ID passado
package getcommonglass

import "github.com/realnfcs/ultividros-project/api/domain/repository"

// Usecase responsável pela obtenção um único vidro comum de um
// repositório voltando um ponteiro da entidade com o mesmo ID passado
// no parâmetro
type GetCommonGlass struct {
	CommonGlassRepository repository.CommonGlassRepository
}

func (g *GetCommonGlass) Execute(i Input) *Output {
	e, status, err := g.CommonGlassRepository.GetCommonGlass(i.ID)
	return new(Output).Init(e, status, err)
}
