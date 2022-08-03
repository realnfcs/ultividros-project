// Pacote responsável pela o usecase Getsales que executa
// a ação de pegar todas as vendas e os produtos requeridos
// e retornar os dados ao cliente
package getsales

import (
	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável pela obtenção de todas as vendas de
// um repositório voltando um ponteiro de array de vendas
// junto com os produtos requeridos
type GetSales struct {
	SaleRepository repository.SaleRepository
	UserRepository repository.UserRepository
}

func (g *GetSales) Execute(i Input) *Output {
	e, status, err := g.SaleRepository.GetSales(i.ClientId)
	return new(Output).Init(e, status, err)
}
