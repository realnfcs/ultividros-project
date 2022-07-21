package closesale

import (
	"github.com/realnfcs/ultividros-project/api/domain/entities"
)

// Usecase Input Port responsável pelos dados que entrarão
type Input struct {
	Id       string `json:"id"`
	ClientId string `json:"client_id"`
	IsActive bool   `json:"is_active"`
}

// Método que converte um input em uma entidade de venda
func (i *Input) ConvertToSale() *entities.Sale {
	return &entities.Sale{
		Id:       i.Id,
		ClientId: i.ClientId,
		IsActive: i.IsActive,
	}
}
