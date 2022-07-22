package patchpart

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

func (*Input) Init(e entities.Part) *Input {
	return &Input{
		e.Id,
		e.Name,
		e.Description,
		e.Price,
		e.Quantity,
		e.ForType,
	}
}

// Método responsável em converter um input em uma entidade de peça
func (i *Input) ConvertToPart() *entities.Part {
	return &entities.Part{
		Id:          i.Id,
		Name:        i.Name,
		Description: i.Description,
		Price:       i.Price,
		Quantity:    i.Quantity,
		ForType:     i.ForType,
	}
}
