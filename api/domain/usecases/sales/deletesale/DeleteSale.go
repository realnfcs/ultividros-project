// Pacote responsável pela o usecase DeleteSale que executa
// a ação de deletar uma venda no repositório
package deletesale

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável por deletar uma venda no repositório
type DeleteSale struct {
	SaleRepository repository.SaleRepository
}

func (d *DeleteSale) Execute(i Input) *Output {
	// Verificação se a venda já foi finalizada (isActive == false)
	// ou saber se os produtos foram todos cancelados

	if i.IsActive {
		return new(Output).Init(400, errors.New("can't delete a unfinished sale"))
	}

	status, err := d.SaleRepository.DeleteSale(*i.ConvertToSale())
	return new(Output).Init(status, err)
}