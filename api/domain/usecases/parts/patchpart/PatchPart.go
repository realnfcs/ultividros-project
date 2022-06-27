// Pacote responsável pela o usecase PatchPart que executa
// a ação de atualizar somente os campos que tiveram mudança de uma
// peça salvando-os no repositório
package patchpart

import "github.com/realnfcs/ultividros-project/api/domain/repository"

type PatchPart struct {
	PartRepository repository.PartRepository
}

func (p *PatchPart) Execute(i Input) *Output {
	id, status, err := p.PartRepository.PatchPart(*i.ConvertToPart())
	return new(Output).Init(id, status, err)
}
