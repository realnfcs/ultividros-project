// Pacote responsável pela o usecase GetCommonGlasses que executa
// a ação de pegar todos os vidro comuns e retornar os dados ao
// cliente
package getcommonglasses

import "github.com/realnfcs/ultividros-project/api/domain/repository"

// Usecase responsável pela obtenção de todos os vidros comuns de um
// repositório voltando um ponteiro de array de vidros comuns
type GetCommonGlasses struct {
	CommonGlassRepository repository.CommonGlassRepository
}

// Método que executa o procedimento de pegar os vidros comuns
func (g *GetCommonGlasses) Execute() *Output {
	e, status, err := g.CommonGlassRepository.GetCommonGlasses()
	return new(Output).Init(e, status, err)
}
