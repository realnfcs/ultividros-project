// Pacote responsável pela o usecase GetTemperedGlass que executa
// a ação de deletar um vidro temperado no repositório
package deletetemperedglass

import "github.com/realnfcs/ultividros-project/api/domain/repository"

// Usecase responsável por deletar um vidro temperado no repositório
type DeleteTemperedGlass struct {
	TemperedGlassRepository repository.TemperedGlassRepository
}

func (d *DeleteTemperedGlass) Execute(i Input) *Output {
	status, err := d.TemperedGlassRepository.DeleteTemperedGlass(*i.ConvertToTempGlss())
	return new(Output).Init(status, err)
}
