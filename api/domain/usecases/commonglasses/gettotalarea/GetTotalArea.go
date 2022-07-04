// Pacote responsável pela o usecase GetTotalArea que executa
// a ação de pegar a altura e comprimento de um vidro comum e
// retornar os dados
package gettotalarea

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável pela obtenção da altura e comprimento de
// um vidro comum de um repositório voltando um map com os dados
type GetTotalArea struct {
	CommonGlassRepository repository.CommonGlassRepository
}

func (g *GetTotalArea) Execute(i Input) *Output {
	if i.Id == "" {
		return new(Output).Init(i.Id, 0, 0, errors.New("No id error"))
	}

	area, err := g.CommonGlassRepository.GetArea(i.Id)
	return new(Output).Init(i.Id, area["width"], area["height"], err)
}
