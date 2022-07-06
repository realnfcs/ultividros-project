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

// Essas structs provém as informações base contidas na entidade de vendas
// porém com mais enfâse nas bibliotecas usadas.
type Sale struct {
	ID          string       `json:"id" gorm:"primaryKey"`
	ClientID    string       `json:"client_id"`
	ProductsIDs []string     `json:"products_ids"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at" gorm:"index"`
}

// Método para criar um uuid antes de salvar no banco de dados
func (m *Sale) BeforeCreate(scope *gorm.DB) error {
	id := uuid.New().String()
	if id == "" {
		return errors.New("Cannot create uuid")
	}

	m.ID = strings.Replace(id, "-", "", -1)

	return nil
}

// Método responsável por transformar um Model em uma entidades Sale
func (m *Sale) TranformToEntity() *entities.Sale {
	return &entities.Sale{
		Id:          m.ID,
		ClientId:    m.ClientID,
		ProductsIds: m.ProductsIDs,
	}
}

// Método responsável por transformar um Slice de Models em um Slice de entidades Sale
func (*Sale) TransformToSliceOfEntity(m []Sale) *[]entities.Sale {

	sale := make([]entities.Sale, len(m))

	var wg sync.WaitGroup

	channel := make(chan int)

	channel <- 0

	for range m {

		wg.Add(1)

		go func() {

			sale[<-channel].Id = m[<-channel].ID
			sale[<-channel].ClientId = m[<-channel].ClientID
			sale[<-channel].ProductsIds = m[<-channel].ProductsIDs

			channel <- <-channel + 1

			wg.Done()
		}()

		if index := <-channel + 1; len(m)-1 > index {

			wg.Add(1)

			go func() {
				sale[<-channel+1].Id = m[<-channel+1].ID
				sale[<-channel+1].ClientId = m[<-channel+1].ClientID
				sale[<-channel+1].ProductsIds = m[<-channel+1].ProductsIDs

				channel <- <-channel + 1

				wg.Done()
			}()
		}

		wg.Wait()
	}

	close(channel)

	return &sale

}

// Método responsável por transformar a entidade Sale em model
func (m *Sale) TransformToModel(e entities.Sale) *Sale {
	return &Sale{
		e.Id,
		e.ClientId,
		e.ProductsIds,
		time.Time{},
		time.Time{},
		sql.NullTime{},
	}
}

// Método que transfoma um Slice de entidades em Slice de models Sales
func (*Sale) TransformToSliceOfModel(e []entities.Sale) *[]Sale {

	var (
		m  []Sale
		wg sync.WaitGroup
	)

	channel := make(chan int)

	channel <- 0

	for range e {

		wg.Add(1)

		go func() {

			m[<-channel].ID = e[<-channel].Id
			m[<-channel].ClientID = e[<-channel].ClientId
			m[<-channel].ProductsIDs = e[<-channel].ProductsIds

			channel <- <-channel + 1

			wg.Done()
		}()

		if index := <-channel + 1; len(e)-1 > index {

			wg.Add(1)

			go func() {
				m[<-channel+1].ID = e[<-channel+1].Id
				m[<-channel+1].ClientID = e[<-channel+1].ClientId
				m[<-channel+1].ProductsIDs = e[<-channel+1].ProductsIds

				channel <- <-channel + 1

				wg.Done()
			}()
		}

		wg.Wait()
	}

	close(channel)

	return &m
}
