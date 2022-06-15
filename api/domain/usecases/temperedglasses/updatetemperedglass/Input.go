package updatetemperedglass

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// Usecase Input Port responsável pelos dados que entrarão
type Input struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    uint32  `json:"quantity"`
	Type        string  `json:"type"`
	Color       string  `json:"color"`
	GlassSheets uint8   `json:"glass_sheets"`
	Milimeter   float32 `json:"milimeter"`
	Height      float32 `json:"height"`
	Width       float32 `json:"width"`
}

func (*Input) Init(e entities.TemperedGlass) *Input {
	return &Input{
		e.Id,
		e.Name,
		e.Description,
		e.Price,
		e.Quantity,
		e.Type,
		e.Color,
		e.GlassSheets,
		e.Milimeter,
		e.Height,
		e.Width,
	}
}

// Método responsável em converter um input em uma entidade de vidro temperado
func (i *Input) ConvertToTempGlss() *entities.TemperedGlass {
	return &entities.TemperedGlass{
		i.ID,
		i.Name,
		i.Description,
		i.Price,
		i.Quantity,
		i.Type,
		i.Color,
		i.GlassSheets,
		i.Milimeter,
		i.Height,
		i.Width,
	}
}
