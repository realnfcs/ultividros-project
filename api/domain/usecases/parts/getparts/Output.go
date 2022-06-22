package getparts

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// Usecase Output Port responsável pelos dados que serão retornados
type Output struct {
	Data   []OutputData `json:"data"`
	Status int          `json:"-"`
	Err    string       `json:"error"`
}

// OutputData é a estrutura de dados que será retornado em um array
// no Output
type OutputData struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    uint32  `json:"quantity"`
	ForType     string  `json:"for_type"`
}

func (*Output) Init(i *[]entities.Part, status int, err error) *Output {

	output := make([]OutputData, len(*i))

	if i == nil {
		return &Output{output, status, err.Error()}
	}

	if err != nil {
		return &Output{output, status, err.Error()}
	}

	for i, v := range *i {
		output[i].Id = v.Id
		output[i].Name = v.Name
		output[i].Description = v.Description
		output[i].Price = v.Price
		output[i].Quantity = v.Quantity
		output[i].ForType = v.ForType
	}

	return &Output{output, status, ""}
}
