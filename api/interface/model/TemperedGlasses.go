// Os models, na parte de Interface and Adapters, vai servir como uma
// estrutura de dados que fica entre as entidades, os viewmodels e as
// tabelas do banco de dados. Por exemplo, nele terá a biblioteca time
// do Go para haver uma facilidade no gerenciamento de datas.
package model

import "github.com/NicolasSales0101/ultividros-project/api/domain/entities"

// Essa struct provém as informações base contidas na entidade de vidros
// temperados porém com mais enfâse nas bibliotecas usadas.

// TODO: Criar os Input e Output do DTO que estamos criando e implementa-los
// nos arquivos dentro da pasta interface que representa a camada Interface
// and Adapters
type ModelTemperedGlass struct {
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

func (m *ModelTemperedGlass) Init(e entities.TemperedGlass) *ModelTemperedGlass {
	return &ModelTemperedGlass{
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

func (*ModelTemperedGlass) Map(e []entities.TemperedGlass) *[]ModelTemperedGlass {
	var m []ModelTemperedGlass

	for i, v := range e {
		go func() {
			m[i].Id = v.Id
			m[i].Name = v.Name
			m[i].Description = v.Description
			m[i].Price = v.Price
			m[i].Quantity = v.Quantity
			m[i].Type = v.Type
			m[i].Color = v.Color
			m[i].GlassSheets = v.GlassSheets
			m[i].Milimeter = v.Milimeter
			m[i].Height = v.Height
			m[i].Width = v.Width
		}()

		if index := i + 1; len(e) > index {
			go func() {
				m[i+1].Id = e[i+1].Id
				m[i+1].Name = e[i+1].Name
				m[i+1].Description = e[i+1].Description
				m[i+1].Price = e[i+1].Price
				m[i+1].Quantity = e[i+1].Quantity
				m[i+1].Type = e[i+1].Type
				m[i+1].Color = e[i+1].Color
				m[i+1].GlassSheets = e[i+1].GlassSheets
				m[i+1].Milimeter = e[i+1].Milimeter
				m[i+1].Height = e[i+1].Height
				m[i+1].Width = e[i+1].Width
			}()
		}
	}

	return &m
}
