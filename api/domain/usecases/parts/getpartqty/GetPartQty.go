// Pacote responsável pela o usecase GetPartQuantity que executa a
// ação de pegar a quantidade total de uma peça e retornar os dados
package getpartqty

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável pela obtenção da quantidade total de
// uma peça de um repositório
type GetPartQty struct {
	PartRepository repository.PartRepository
}

func (g *GetPartQty) Execute(i Input) *Output {
	if i.Id == "" {
		return new(Output).Init(i.Id, 0, errors.New("No id error"))
	}

	qty, err := g.PartRepository.GetPartQuantity(i.Id)
	return new(Output).Init(i.Id, qty, err)
}
