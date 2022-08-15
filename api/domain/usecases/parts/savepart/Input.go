package savepart

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// Usecase Input Port responsável pelos dados que entrarão
type Input struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    uint32  `json:"quantity"`
	ForType     string  `json:"for_type"`
	UserId      string  `json:"user_id"`
}

func (*Input) Init(e entities.Part) *Input {
	return &Input{
		e.Name,
		e.Description,
		e.Price,
		e.Quantity,
		e.ForType,
		"",
	}
}

// Método que converte um input na entidade Part
func (i *Input) ConvertToPart() *entities.Part {
	return &entities.Part{
		Id:          "",
		Name:        i.Name,
		Description: i.Description,
		Price:       i.Price,
		Quantity:    i.Quantity,
		ForType:     i.ForType,
	}
}
