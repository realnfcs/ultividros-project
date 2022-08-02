package deletetemperedglass

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
	UserId      string  `json:"user_id"`
}

// Método que converte um input em uma entidade de vidro temperado
func (i *Input) ConvertToTempGlss() *entities.TemperedGlass {
	return &entities.TemperedGlass{
		Id:          i.ID,
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
