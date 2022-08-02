// Pacote responsável pela o usecase DeletePart que executa
// a ação de deletar uma peça no repositório
package deletepart

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável por deletar uma peça no repositório
type DeletePart struct {
	PartRepository repository.PartRepository
	UserRepository repository.UserRepository
}

func (d *DeletePart) Execute(i Input) *Output {

	ocup, err := d.UserRepository.VerifyOccupation(i.UserId)
	if err != nil {
		return new(Output).Init(400, err)
	}

	if !ocup {
		return new(Output).Init(401, errors.New("unauthorized"))
	}

	status, err := d.PartRepository.DeletePart(*i.ConvertToPart())
	return new(Output).Init(status, err)
}
