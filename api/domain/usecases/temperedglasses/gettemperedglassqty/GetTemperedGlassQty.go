// Pacote responsável pela o usecase GetTemperedGlassQty que executa
// a ação de pegar a quantidade total de um vidro temperado e
// retornar os dados
package gettemperedglassqty

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável pela obtenção da quantidade total de
// um vidro temperado de um repositório
type GetQuantity struct {
	TemperedGlassRepository repository.TemperedGlassRepository
}

func (g *GetQuantity) Execute(i Input) *Output {
	if i.Id == "" {
		return new(Output).Init(i.Id, 0, errors.New("No id error"))
	}

	qty, err := g.TemperedGlassRepository.GetTempGlssQty(i.Id)
	return new(Output).Init(i.Id, qty, err)
}
