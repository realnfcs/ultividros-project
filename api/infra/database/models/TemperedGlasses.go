// Os models, na parte de Interface and Adapters, vai servir como uma
// estrutura de dados que fica entre as entidades, os viewmodels e as
// tabelas do banco de dados. Por exemplo, nele terá a biblioteca time
// do Go para haver uma facilidade no gerenciamento de datas.
package models

import (
	"database/sql"
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/realnfcs/ultividros-project/api/domain/entities"
	"gorm.io/gorm"
)

// Essa struct provém as informações base contidas na entidade de vidros
// temperados porém com mais enfâse nas bibliotecas usadas.
type ModelTemperedGlass struct {
	ID          string       `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Price       float32      `json:"price"`
	Quantity    uint32       `json:"quantity"`
	Type        string       `json:"type"`
	Color       string       `json:"color"`
	GlassSheets uint8        `json:"glass_sheets"`
	Milimeter   float32      `json:"milimeter"`
	Height      float32      `json:"height"`
	Width       float32      `json:"width"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at" gorm:"index"`
}

// Método para criar um uuid antes de salvar no banco de dados
func (m *ModelTemperedGlass) BeforeCreate(scope *gorm.DB) error {
	id := uuid.New().String()
	if id == "" {
		return errors.New("Cannot create uuid")
	}

	m.ID = strings.Replace(id, "-", "", -1)
	// scope.Statement.SetColumn("ID", uuid)

	return nil
}

// Método responsável por transformar o model em entidade
func (m *ModelTemperedGlass) TranformToEntity() *entities.TemperedGlass {
	return &entities.TemperedGlass{
		Id:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
		Quantity:    m.Quantity,
		Type:        m.Type,
		Color:       m.Color,
		GlassSheets: m.GlassSheets,
		Milimeter:   m.Milimeter,
		Height:      m.Height,
		Width:       m.Width,
	}
}

// Método responsável por transformar um Slice de Models em um Slice de entidades
func (*ModelTemperedGlass) TranformToSliceOfEntity(m []ModelTemperedGlass) *[]entities.TemperedGlass {

	tempGlasses := make([]entities.TemperedGlass, len(m))

	var wg sync.WaitGroup

	for i, v := range m {

		wg.Add(1)

		go func() {
			tempGlasses[i].Id = v.ID
			tempGlasses[i].Name = v.Name
			tempGlasses[i].Description = v.Description
			tempGlasses[i].Price = v.Price
			tempGlasses[i].Quantity = v.Quantity
			tempGlasses[i].Type = v.Type
			tempGlasses[i].Color = v.Color
			tempGlasses[i].GlassSheets = v.GlassSheets
			tempGlasses[i].Milimeter = v.Milimeter
			tempGlasses[i].Height = v.Height
			tempGlasses[i].Width = v.Width

			wg.Done()
		}()

		if index := i + 1; len(m)-1 > index {

			wg.Add(1)

			go func() {
				tempGlasses[i+1].Id = m[i+1].ID
				tempGlasses[i+1].Name = m[i+1].Name
				tempGlasses[i+1].Description = m[i+1].Description
				tempGlasses[i+1].Price = m[i+1].Price
				tempGlasses[i+1].Quantity = m[i+1].Quantity
				tempGlasses[i+1].Type = m[i+1].Type
				tempGlasses[i+1].Color = m[i+1].Color
				tempGlasses[i+1].GlassSheets = m[i+1].GlassSheets
				tempGlasses[i+1].Milimeter = m[i+1].Milimeter
				tempGlasses[i+1].Height = m[i+1].Height
				tempGlasses[i+1].Width = m[i+1].Width

				wg.Done()
			}()
		}

		wg.Wait()
	}

	return &tempGlasses
}

// Método que transforma uma entidade em model
func (m *ModelTemperedGlass) TransformToModel(e entities.TemperedGlass) *ModelTemperedGlass {
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
		time.Time{},
		time.Time{},
		sql.NullTime{},
	}
}

// Método que transfoma um Slice de entidades em Slice de models
func (*ModelTemperedGlass) TransformToSliceOfModel(e []entities.TemperedGlass) *[]ModelTemperedGlass {
	var m []ModelTemperedGlass

	var wg sync.WaitGroup

	for i, v := range e {

		wg.Add(1)

		go func() {
			m[i].ID = v.Id
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

			wg.Done()
		}()

		if index := i + 1; len(e)-1 > index {

			wg.Add(1)

			go func() {
				m[i+1].ID = e[i+1].Id
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

				wg.Done()
			}()
		}

		wg.Wait()
	}

	return &m
}
