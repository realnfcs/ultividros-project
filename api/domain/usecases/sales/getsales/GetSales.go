// Pacote responsável pela o usecase Getsales que executa
// a ação de pegar todas as vendas e os produtos requeridos
// e retornar os dados ao cliente
package getsales

import "github.com/realnfcs/ultividros-project/api/domain/repository"

// Usecase responsável pela obtenção de todas as vendas de
// um repositório voltando um ponteiro de array de vendas
// junto com os produtos requeridos
type GetSales struct {
	SaleRepository repository.SaleRepository
}

func (g *GetSales) Execute() *Output {
	e, status, err := g.SaleRepository.GetSales()
	return new(Output).Init(e, status, err)
}
