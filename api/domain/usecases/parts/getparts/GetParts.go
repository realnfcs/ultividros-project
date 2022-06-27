// Pacote responsável pela o usecase GetParts que executa
// a ação de pegar todas as peças e retornar os dados ao
// cliente
package getparts

import "github.com/realnfcs/ultividros-project/api/domain/repository"

// Usecase responsável pela obtenção de todas peças de um
// repositório voltando um ponteiro de array de peças
type GetParts struct {
	PartRepository repository.PartRepository
}

// Método que executa o procedimento de pegar as peças
func (g *GetParts) Execute() *Output {
	e, status, err := g.PartRepository.GetParts()
	return new(Output).Init(e, status, err)
}
