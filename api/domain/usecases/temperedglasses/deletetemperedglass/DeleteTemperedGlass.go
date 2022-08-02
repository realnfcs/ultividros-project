// Pacote responsável pela o usecase DeleteTemperedGlass que executa
// a ação de deletar um vidro temperado no repositório
package deletetemperedglass

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável por deletar um vidro temperado no repositório
type DeleteTemperedGlass struct {
	TemperedGlassRepository repository.TemperedGlassRepository
	UserRepository          repository.UserRepository
}

func (d *DeleteTemperedGlass) Execute(i Input) *Output {

	ocup, err := d.UserRepository.VerifyOccupation(i.UserId)
	if err != nil {
		return new(Output).Init(400, err)
	}

	if !ocup {
		return new(Output).Init(401, errors.New("unauthorized"))
	}

	status, err := d.TemperedGlassRepository.DeleteTemperedGlass(*i.ConvertToTempGlss())
	return new(Output).Init(status, err)
}
