package savetemperedglass

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// Usecase Input Port responsável pelos dados que entrarão
type Input struct {
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
	UserId      string  `json:"user_id"`
}

func (*Input) Init(e entities.TemperedGlass) *Input {
	return &Input{
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
		"",
	}
}

// Método que converte um input na entidade CommonGlass
func (i *Input) ConvertToTempGlss() *entities.TemperedGlass {
	return &entities.TemperedGlass{
		Id:          "",
		Name:        i.Name,
		Description: i.Description,
		Price:       i.Price,
		Quantity:    i.Quantity,
		Type:        i.Type,
		Color:       i.Color,
		GlassSheets: i.GlassSheets,
		Milimeter:   i.Milimeter,
		Height:      i.Height,
		Width:       i.Width,
	}
}
