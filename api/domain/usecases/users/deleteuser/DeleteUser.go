// Pacote responsável pela o usecase DeleteUser que executa
// a ação de deletar um usuário no repositório
package deleteuser

import (
	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável por deletar um usuário no repositório
type DeleteUser struct {
	UserRepository repository.UserRepository
}

func (d *DeleteUser) Execute(i Input) *Output {

	status, err := d.UserRepository.DeleteUser(*i.ConvertToUser())
	return new(Output).Init(status, err)
}
