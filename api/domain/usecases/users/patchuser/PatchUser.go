// Pacote responsável pela o usecase PatchUser que executa
// a ação de atualizar somente os campos que tiveram mudança de um
// usuário salvando-os no repositório
package patchuser

import "github.com/realnfcs/ultividros-project/api/domain/repository"

type PatchUser struct {
	UserRepository repository.UserRepository
}

func (p *PatchUser) Execute(i Input) *Output {
	id, status, err := p.UserRepository.PatchUser(*i.ConvertToUser())
	return new(Output).Init(id, status, err)
}
