// Pacote responsável pela o usecase GetSale que executa
// a ação de pegar uma venda e seus produtos requisitados
// e retornar os dados de acordo com o ID passado
package getsale

import "github.com/realnfcs/ultividros-project/api/domain/repository"

type GetSale struct {
	SaleRepository repository.SaleRepository
}

// Usecase responsável pela obtenção de uma única venda de um
// repositório junto com os seus produtos requeridos voltando
// um ponteiro da entidade com o mesmo ID passado no parâmetro
func (g *GetSale) Execute(i Input) *Output {
	e, status, err := g.SaleRepository.GetSale(i.ID)
	return new(Output).Init(e, status, err)
}
