// Pacote responsável pela o usecase DeleteCommonGlass que executa
// a ação de deletar um vidro comum no repositório
package deletecommonglass

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável por deletar um vidro comum no repositório
type DeleteCommonGlass struct {
	CommonGlassRepository repository.CommonGlassRepository
	UserRepository        repository.UserRepository
}

func (d *DeleteCommonGlass) Execute(i Input) *Output {

	ocup, err := d.UserRepository.VerifyOccupation(i.UserId)
	if err != nil {
		return new(Output).Init(400, err)
	}

	if !ocup {
		return new(Output).Init(401, errors.New("unauthorized"))
	}

	status, err := d.CommonGlassRepository.DeleteCommonGlass(*i.ConvertToComnGlss())
	return new(Output).Init(status, err)
}
