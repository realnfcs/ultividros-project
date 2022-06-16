package getcommonglass

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// Usecase Output Port responsável pelos dados que serão retornados
type Output struct {
	OutputData `json:"data"`
	Status     int    `json:"-"`
	Err        string `json:"error"`
}

// OutpurData respnsável pelo dado da entitite em si que será passado pelas camadas externas
type OutputData struct {
	Id              string  `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float32 `json:"price"`
	Type            string  `json:"type"`
	Color           string  `json:"color"`
	Milimeter       float32 `json:"milimeter"`
	HeightAvailable float32 `json:"height_available"`
	WidthAvailable  float32 `json:"width_available"`
}

func (*Output) Init(e *entities.CommonGlass, status int, err error) *Output {
	if e != nil {
		return &Output{
			OutputData{
				e.Id,
				e.Name,
				e.Description,
				e.Price,
				e.Type,
				e.Color,
				e.Milimeter,
				e.HeightAvailable,
				e.WidthAvailable,
			},
			status,
			"",
		}
	}

	return &Output{OutputData{}, status, err.Error()}
}
