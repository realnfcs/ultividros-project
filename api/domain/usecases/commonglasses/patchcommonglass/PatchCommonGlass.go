// Pacote responsável pela o usecase PatchCommonGlass que executa
// a ação de atualizar somente os campos que tiveram mudança de um
// vidro comum salvando-os no repositório
package patchcommonglass

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável por editar um vidro comum no repositório
type PatchCommonGlass struct {
	CommonGlassRepository repository.CommonGlassRepository
	UserRepository        repository.UserRepository
}

func (p *PatchCommonGlass) Execute(i Input) *Output {

	ocup, err := p.UserRepository.VerifyOccupation(i.UserId)
	if err != nil {
		return new(Output).Init(i.ID, 400, err)
	}

	if !ocup {
		return new(Output).Init(i.ID, 401, errors.New("unauthorized"))
	}

	id, status, err := p.CommonGlassRepository.PatchCommonGlass(*i.ConvertToComnGlss())
	return new(Output).Init(id, status, err)
}
