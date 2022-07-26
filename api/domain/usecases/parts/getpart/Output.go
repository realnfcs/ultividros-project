package getpart

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// Usecase Output Port responsável pelos dados que serão retornados
type Output struct {
	OutputData `json:"data"`
	Status     int    `json:"-"`
	Err        string `json:"error"`
}

// OutpurData responsável pelo dado da entity em si que será passado pelas camadas externas
type OutputData struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    uint32  `json:"quantity"`
	ForType     string  `json:"for_type"`
}

func (*Output) Init(e *entities.Part, status int, err error) *Output {
	if e != nil {
		return &Output{
			OutputData{
				e.Id,
				e.Name,
				e.Description,
				e.Price,
				e.Quantity,
				e.ForType,
			},
			status,
			"",
		}
	}

	return &Output{OutputData{}, status, err.Error()}
}
