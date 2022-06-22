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

// Essa struct provém as informações base contidas na entidade de peças
// porém com mais enfâse nas bibliotecas usadas.
type Part struct {
	ID          string       `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Price       float32      `json:"price"`
	Quantity    uint32       `json:"quantity"`
	ForType     string       `json:"for_type"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at" gorm:"index"`
}

// Método para criar um uuid antes de salvar no banco de dados
func (m *Part) BeforeCreate(scope *gorm.DB) error {
	id := uuid.New().String()
	if id == "" {
		return errors.New("Cannot create uuid")
	}

	m.ID = strings.Replace(id, "-", "", -1)

	return nil
}

// Método responsável por transformar o model Part em entidade
func (m *Part) TranformToEntity() *entities.Part {
	return &entities.Part{
		Id:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
		Quantity:    m.Quantity,
		ForType:     m.ForType,
	}
}

// Método responsável por transformar um Slice de Models em um Slice de entidades Part
func (*Part) TranformToSliceOfEntity(m []Part) *[]entities.Part {

	part := make([]entities.Part, len(m))

	var wg sync.WaitGroup

	for i, v := range m {

		wg.Add(1)

		go func() {
			part[i].Id = v.ID
			part[i].Name = v.Name
			part[i].Description = v.Description
			part[i].Price = v.Price
			part[i].Quantity = v.Quantity
			part[i].ForType = v.ForType

			wg.Done()
		}()

		if index := i + 1; len(m)-1 > index {

			wg.Add(1)

			go func() {
				part[i+1].Id = m[i+1].ID
				part[i+1].Name = m[i+1].Name
				part[i+1].Description = m[i+1].Description
				part[i+1].Price = m[i+1].Price
				part[i+1].Quantity = m[i+1].Quantity
				part[i+1].ForType = m[i+1].ForType

				wg.Done()
			}()
		}

		wg.Wait()
	}

	return &part
}

// Método responsável por transformar a entidade Part em model
func (m *Part) TransformToModel(e entities.Part) *Part {
	return &Part{
		e.Id,
		e.Name,
		e.Description,
		e.Price,
		e.Quantity,
		e.ForType,
		time.Time{},
		time.Time{},
		sql.NullTime{},
	}
}

// Método que transfoma um Slice de entidades em Slice de models Parts
func (*Part) TransformToSliceOfModel(e []entities.Part) *[]Part {

	var m []Part

	var wg sync.WaitGroup

	for i, v := range e {

		wg.Add(1)

		go func() {
			m[i].ID = v.Id
			m[i].Name = v.Name
			m[i].Description = v.Description
			m[i].Price = v.Price
			m[i].Quantity = v.Quantity
			m[i].ForType = v.ForType

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
				m[i+1].ForType = e[i+1].ForType

				wg.Done()
			}()
		}

		wg.Wait()
	}

	return &m
}
