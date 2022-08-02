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
	UserId          string  `json:"user_id"`
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
		"",
	}
}

// Método responsável em converter um input em uma entidade de vidro comum
func (i *Input) ConvertToComnGlss() *entities.CommonGlass {
	return &entities.CommonGlass{
		Id:              i.ID,
		Name:            i.Name,
		Description:     i.Description,
		Price:           i.Price,
		Type:            i.Type,
		Color:           i.Color,
		Milimeter:       i.Milimeter,
		HeightAvailable: i.HeightAvailable,
		WidthAvailable:  i.WidthAvailable,
	}
}
