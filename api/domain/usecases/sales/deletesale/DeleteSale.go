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
	UserRepository repository.UserRepository
}

func (d *DeleteSale) Execute(i Input) *Output {
	// Verificação se a venda já foi finalizada (isActive == false)
	// ou saber se os produtos foram todos cancelados

	if i.Id == "" || i.ClientId == "" || (len(i.CommonGlssReq) == 0 && len(i.PartsReq) == 0 && len(i.TempGlssReq) == 0) {
		return new(Output).Init(400, errors.New("some important fields don't have a value"))
	}

	if i.IsActive {
		return new(Output).Init(400, errors.New("can't delete a unfinished sale"))
	}

	clientAuth, err := d.SaleRepository.ClientAuthentication(i.Id, i.ClientId)
	if err != nil {
		return new(Output).Init(500, err)
	}

	if !clientAuth {
		return new(Output).Init(400, errors.New("client unmatched"))
	}

	status, err := d.SaleRepository.DeleteSale(*i.ConvertToSale())
	return new(Output).Init(status, err)
}
