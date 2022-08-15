// Pacote responsável pela o usecase CloseSale que executa
// a ação de finalizar uma venda no repositório caso tudo
// esteja de acordo (a compra e o pedido forem finalizados)
package closesale

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável por finalizar uma venda no repositório
type CloseSale struct {
	SaleRepository repository.SaleRepository
}

func (c *CloseSale) Execute(i Input) *Output {

	if i.Id == "" || i.ClientId == "" {
		return new(Output).Init(i.Id, 400, errors.New("id of sale or client don't have a value"))
	}

	if !i.IsActive {
		return new(Output).Init(i.Id, 400, errors.New("this sale is already closed"))
	}

	clientAuth, err := c.SaleRepository.ClientAuthentication(i.Id, i.ClientId)
	if err != nil {
		return new(Output).Init(i.Id, 500, err)
	}

	if !clientAuth {
		return new(Output).Init(i.Id, 400, errors.New("client unmatched"))
	}

	id, status, err := c.SaleRepository.CloseSale(*i.ConvertToSale())
	return new(Output).Init(id, status, err)
}
