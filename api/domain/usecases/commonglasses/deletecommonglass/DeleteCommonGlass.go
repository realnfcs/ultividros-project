// Pacote responsável pela o usecase DeleteCommonGlass que executa
// a ação de deletar um vidro comum no repositório
package deletecommonglass

import "github.com/realnfcs/ultividros-project/api/domain/repository"

// Usecase responsável por deletar um vidro temperado no repositório
type DeleteCommonGlass struct {
	CommonGlassRepository repository.CommonGlassRepository
}

func (d *DeleteCommonGlass) Execute(i Input) *Output {
	status, err := d.CommonGlassRepository.DeleteCommonGlass(*i.ConvertToComnGlss())
	return new(Output).Init(status, err)
}
