// Pacote responsável pela o usecase PatchPart que executa
// a ação de atualizar somente os campos que tiveram mudança de uma
// peça salvando-os no repositório
package patchpart

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável em editar uma peça no repositório
type PatchPart struct {
	PartRepository repository.PartRepository
	UserRepository repository.UserRepository
}

func (p *PatchPart) Execute(i Input) *Output {

	ocup, err := p.UserRepository.VerifyOccupation(i.UserId)
	if err != nil {
		return new(Output).Init("", 400, err)
	}

	if !ocup {
		return new(Output).Init("", 401, errors.New("unauthorized"))
	}

	id, status, err := p.PartRepository.PatchPart(*i.ConvertToPart())
	return new(Output).Init(id, status, err)
}
