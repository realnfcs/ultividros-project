package deletepart

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// Usecase Input Port responsável pelos dados que entrarão
type Input struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    uint32  `json:"quantity"`
	ForType     string  `json:"for_type"`
}

// Método que converte um input em uma entidade de peça
func (i *Input) ConvertToPart() *entities.Part {
	return &entities.Part{
		i.Id,
		i.Name,
		i.Description,
		i.Price,
		i.Quantity,
		i.ForType,
	}
}
