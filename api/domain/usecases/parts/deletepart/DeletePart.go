// Pacote responsável pela o usecase DeletePart que executa
// a ação de deletar uma peça no repositório
package deletepart

import "github.com/realnfcs/ultividros-project/api/domain/repository"

// Usecase responsável por deletar uma peça no repositório
type DeletePart struct {
	PartRepository repository.PartRepository
}

func (d *DeletePart) Execute(i Input) *Output {
	status, err := d.PartRepository.DeletePart(*i.ConvertToPart())
	return new(Output).Init(status, err)
}
