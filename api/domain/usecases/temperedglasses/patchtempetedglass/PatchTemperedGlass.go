// Pacote responsável pela o usecase PatchTemperedGlass que executa
// a ação de atualizar somente os campos que tiveram mudança de um
// vidro temperado salvando-os no repositório
package patchtemperedglass

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável em editar um vidro temperado no repositório
type PatchTemperedGlass struct {
	TemperedGlassRepository repository.TemperedGlassRepository
	UserRepository          repository.UserRepository
}

func (p *PatchTemperedGlass) Execute(i Input) *Output {

	ocup, err := p.UserRepository.VerifyOccupation(i.UserId)
	if err != nil {
		return new(Output).Init("", 400, err)
	}

	if !ocup {
		return new(Output).Init("", 401, errors.New("unauthorized"))
	}

	id, status, err := p.TemperedGlassRepository.PatchTemperedGlass(*i.ConvertToTempGlss())
	return new(Output).Init(id, status, err)
}
