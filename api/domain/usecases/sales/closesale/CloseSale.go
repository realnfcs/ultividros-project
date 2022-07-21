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

	if !i.IsActive {
		return new(Output).Init(i.Id, 400, errors.New("this sale is already closed"))
	}

	id, status, err := c.SaleRepository.CloseSale(*i.ConvertToSale())
	return new(Output).Init(id, status, err)
}
