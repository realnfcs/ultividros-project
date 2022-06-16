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
// comuns porém com mais enfâse nas bibliotecas usadas.
type CommonGlass struct {
	ID              string       `json:"id" gorm:"primaryKey"`
	Name            string       `json:"name"`
	Description     string       `json:"description"`
	Price           float32      `json:"price"`
	Type            string       `json:"type"`
	Color           string       `json:"color"`
	Milimeter       float32      `json:"milimeter"`
	HeightAvailable float32      `json:"height_available"`
	WidthAvailable  float32      `json:"width_available"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	DeletedAt       sql.NullTime `json:"deleted_at" gorm:"index"`
}

// Método para criar um uuid antes de salvar no banco de dados
func (m *CommonGlass) BeforeCreate(scope *gorm.DB) error {
	id := uuid.New().String()
	if id == "" {
		return errors.New("Cannot create uuid")
	}

	m.ID = strings.Replace(id, "-", "", -1)

	return nil
}

// Método responsável por transformar o model CommonGlass em entidade
func (m *CommonGlass) TranformToEntity() *entities.CommonGlass {
	return &entities.CommonGlass{
		Id:              m.ID,
		Name:            m.Name,
		Description:     m.Description,
		Price:           m.Price,
		Type:            m.Type,
		Color:           m.Color,
		Milimeter:       m.Milimeter,
		HeightAvailable: m.HeightAvailable,
		WidthAvailable:  m.WidthAvailable,
	}
}

// Método responsável por transformar um Slice de Models em um Slice de entidades CommonGlasses
func (*CommonGlass) TranformToSliceOfEntity(m []CommonGlass) *[]entities.CommonGlass {

	comnGlasses := make([]entities.CommonGlass, len(m))

	var wg sync.WaitGroup

	for i, v := range m {

		wg.Add(1)

		go func() {
			comnGlasses[i].Id = v.ID
			comnGlasses[i].Name = v.Name
			comnGlasses[i].Description = v.Description
			comnGlasses[i].Price = v.Price
			comnGlasses[i].Type = v.Type
			comnGlasses[i].Color = v.Color
			comnGlasses[i].Milimeter = v.Milimeter
			comnGlasses[i].HeightAvailable = v.HeightAvailable
			comnGlasses[i].WidthAvailable = v.WidthAvailable

			wg.Done()
		}()

		if index := i + 1; len(m)-1 > index {

			wg.Add(1)

			go func() {
				comnGlasses[i+1].Id = m[i+1].ID
				comnGlasses[i+1].Name = m[i+1].Name
				comnGlasses[i+1].Description = m[i+1].Description
				comnGlasses[i+1].Price = m[i+1].Price
				comnGlasses[i+1].Type = m[i+1].Type
				comnGlasses[i+1].Color = m[i+1].Color
				comnGlasses[i+1].Milimeter = m[i+1].Milimeter
				comnGlasses[i+1].HeightAvailable = m[i+1].HeightAvailable
				comnGlasses[i+1].WidthAvailable = m[i+1].WidthAvailable

				wg.Done()
			}()
		}

		wg.Wait()
	}

	return &comnGlasses
}

// Método responsável por transformar a entidade CommonGlass em model
func (m *CommonGlass) TransformToModel(e entities.CommonGlass) *CommonGlass {
	return &CommonGlass{
		e.Id,
		e.Name,
		e.Description,
		e.Price,
		e.Type,
		e.Color,
		e.Milimeter,
		e.HeightAvailable,
		e.WidthAvailable,
		time.Time{},
		time.Time{},
		sql.NullTime{},
	}
}

// Método que transfoma um Slice de entidades em Slice de models CommonGlasses
func (*CommonGlass) TransformToSliceOfModel(e []entities.CommonGlass) *[]CommonGlass {
	var m []CommonGlass

	var wg sync.WaitGroup

	for i, v := range e {

		wg.Add(1)

		go func() {
			m[i].ID = v.Id
			m[i].Name = v.Name
			m[i].Description = v.Description
			m[i].Price = v.Price
			m[i].Type = v.Type
			m[i].Color = v.Color
			m[i].Milimeter = v.Milimeter
			m[i].HeightAvailable = v.HeightAvailable
			m[i].WidthAvailable = v.WidthAvailable

			wg.Done()
		}()

		if index := i + 1; len(e)-1 > index {

			wg.Add(1)

			go func() {
				m[i+1].ID = e[i+1].Id
				m[i+1].Name = e[i+1].Name
				m[i+1].Description = e[i+1].Description
				m[i+1].Price = e[i+1].Price
				m[i+1].Type = e[i+1].Type
				m[i+1].Color = e[i+1].Color
				m[i+1].Milimeter = e[i+1].Milimeter
				m[i+1].HeightAvailable = e[i+1].HeightAvailable
				m[i+1].WidthAvailable = e[i+1].WidthAvailable

				wg.Done()
			}()
		}

		wg.Wait()
	}

	return &m
}
