// Pacote responsável pela o usecase GetTemperedGlass que executa
// a ação de pegar um vidro temperado e retornar os dados de acordo
// com o ID passado
package gettemperedglass

import (
	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável pela obtenção um único vidro temperado de um
// repositório voltando um ponteiro da entidade com o mesmo ID passado
// no parâmetro
type GetTemperedGlass struct {
	TemperedGlassRepository repository.TemperedGlassRepository
}

func (g *GetTemperedGlass) Execute(i Input) *Output {
	e, status, err := g.TemperedGlassRepository.GetTemperedGlass(i.ID)
	return new(Output).Init(e, status, err)
}
