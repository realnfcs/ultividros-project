package gettemperedglass

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// Usecase Output Port responsável pelos dados que serão retornados
type Output struct {
	OutputData `json:"data"`
	Status     int    `json:"-"`
	Err        string `json:"error"`
}

type OutputData struct {
	Id          string  `json:"id"`
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

func (*Output) Init(e *entities.TemperedGlass, status int, err error) *Output {
	if e != nil {
		return &Output{
			OutputData{
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
			},
			status,
			"",
		}
	}

	return &Output{OutputData{}, status, err.Error()}
}
