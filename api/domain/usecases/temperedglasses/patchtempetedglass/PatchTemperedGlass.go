// Pacote responsável pela o usecase PatchTemperedGlass que executa
// a ação de atualizar somente os campos que tiveram mudança de um
// vidro temperado salvando-os no repositório
package patchtemperedglass

import "github.com/realnfcs/ultividros-project/api/domain/repository"

type PatchTemperedGlass struct {
	TemperedGlassRepository repository.TemperedGlassRepository
}

func (p *PatchTemperedGlass) Execute(i Input) *Output {
	id, status, err := p.TemperedGlassRepository.PatchTemperedGlass(*i.ConvertToTempGlss())
	return new(Output).Init(id, status, err)
}
