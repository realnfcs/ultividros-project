// Pacote responsável pela o usecase GetTemperedGlasses que executa
// a ação de pegar todos os vidro temperados e retornar os dados ao
// cliente
package gettemperedglasses

import (
	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável pela obtenção de todos os vidros temperados de um
// repositório voltando um ponteiro de array de vidros temperados
type GetTemperedGlasses struct {
	TemperedGlassRepository repository.TemperedGlassRepository
}

// Método que executa o procedimento de pegar os vidros temperados
func (g *GetTemperedGlasses) Execute() *Output {
	e, status, err := g.TemperedGlassRepository.GetTemperedGlasses()
	return new(Output).Init(e, status, err)
}
