package patchcommonglass

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// Usecase Input Port responsável pelos dados que entrarão
type Input struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float32 `json:"price"`
	Type            string  `json:"type"`
	Color           string  `json:"color"`
	Milimeter       float32 `json:"milimeter"`
	HeightAvailable float32 `json:"height_available"`
	WidthAvailable  float32 `json:"width_available"`
}

func (*Input) Init(e entities.CommonGlass) *Input {
	return &Input{
		e.Id,
		e.Name,
		e.Description,
		e.Price,
		e.Type,
		e.Color,
		e.Milimeter,
		e.HeightAvailable,
		e.WidthAvailable,
	}
}

// Método responsável em converter um input em uma entidade de vidro comum
func (i *Input) ConvertToComnGlss() *entities.CommonGlass {
	return &entities.CommonGlass{
		i.ID,
		i.Name,
		i.Description,
		i.Price,
		i.Type,
		i.Color,
		i.Milimeter,
		i.HeightAvailable,
		i.WidthAvailable,
	}
}
