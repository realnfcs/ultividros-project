// Pacote responsável pela o usecase PatchCommonGlass que executa
// a ação de atualizar somente os campos que tiveram mudança de um
// vidro comum salvando-os no repositório
package patchcommonglass

import "github.com/realnfcs/ultividros-project/api/domain/repository"

type PatchCommonGlass struct {
	CommonGlassRepository repository.CommonGlassRepository
}

func (p *PatchCommonGlass) Execute(i Input) *Output {
	id, status, err := p.CommonGlassRepository.PatchCommonGlass(*i.ConvertToComnGlss())
	return new(Output).Init(id, status, err)
}
