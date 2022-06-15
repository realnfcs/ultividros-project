// Pacote responsável pela o usecase GetTemperedGlasses que executa
// a ação de pegar todos os vidro temperados e retornar os dados ao
// cliente
package gettemperedglasses

import (
	"github.com/realnfcs/ultividros-project/api/domain/entities"
)

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
	Type        string  `json:"type"`
	Color       string  `json:"color"`
	GlassSheets uint8   `json:"glass_sheets"`
	Milimeter   float32 `json:"milimeter"`
	Height      float32 `json:"height"`
	Width       float32 `json:"width"`
}

func (*Output) Init(i *[]entities.TemperedGlass, status int, err error) *Output {

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
		output[i].Type = v.Type
		output[i].Color = v.Color
		output[i].GlassSheets = v.GlassSheets
		output[i].Milimeter = v.Milimeter
		output[i].Height = v.Height
		output[i].Width = v.Width
	}

	return &Output{output, status, ""}
}
