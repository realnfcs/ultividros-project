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

// Essas structs provém as informações base contidas na entidade de requições
// peças porém com mais enfâse nas bibliotecas usadas. Essa struct terá as
// informações da peça comprada pelo cliente tendo uma relação com o model de
// vendas
type PartReq struct {
	ProductsRequest
}

// Método para criar um uuid antes de salvar no banco de dados
func (m *PartReq) BeforeCreate(scope *gorm.DB) error {
	id := uuid.New().String()
	if id == "" {
		return errors.New("Cannot create uuid")
	}

	m.ID = strings.Replace(id, "-", "", -1)

	return nil
}

// Método responsável por transformar um Model em uma entidades PartsReq
func (m *PartReq) TranformToEntity() *entities.PartsReq {
	return &entities.PartsReq{
		ProductsRequest: entities.ProductsRequest{
			Id:           m.ID,
			ProductId:    m.ProductId,
			ProductName:  m.ProductName,
			ProductPrice: m.ProductPrice,
			ProdtQtyReq:  m.ProdtQtyReq,
			SaleId:       m.SaleID,
		},
	}
}

// Método responsável por transformar um Slice de Models em um Slice de entidades PartReq
func (*PartReq) TransformToSliceOfEntity(m []PartReq) *[]entities.PartsReq {

	part := make([]entities.PartsReq, len(m))

	var wg sync.WaitGroup

	channel := make(chan int)

	channel <- 0

	for range m {

		wg.Add(1)

		go func() {

			part[<-channel].Id = m[<-channel].ID
			part[<-channel].ProductId = m[<-channel].ProductId
			part[<-channel].ProductName = m[<-channel].ProductName
			part[<-channel].ProductPrice = m[<-channel].ProductPrice
			part[<-channel].ProdtQtyReq = m[<-channel].ProdtQtyReq
			part[<-channel].SaleId = m[<-channel].SaleID

			channel <- <-channel + 1

			wg.Done()
		}()

		if index := <-channel + 1; len(m)-1 > index {

			wg.Add(1)

			go func() {
				part[<-channel+1].Id = m[<-channel+1].ID
				part[<-channel+1].ProductId = m[<-channel+1].ProductId
				part[<-channel+1].ProductName = m[<-channel+1].ProductName
				part[<-channel+1].ProductPrice = m[<-channel+1].ProductPrice
				part[<-channel+1].ProdtQtyReq = m[<-channel+1].ProdtQtyReq
				part[<-channel+1].SaleId = m[<-channel+1].SaleID

				channel <- <-channel + 1

				wg.Done()
			}()
		}

		wg.Wait()
	}

	close(channel)

	return &part
}

// Método responsável por transformar a entidade PartsReq em model
func (m *PartReq) TransformToModel(e entities.PartsReq) *PartReq {
	return &PartReq{
		ProductsRequest{
			e.Id,
			e.ProductId,
			e.ProductName,
			e.ProductPrice,
			e.ProdtQtyReq,
			e.SaleId,
			time.Time{},
			time.Time{},
			sql.NullTime{},
		},
	}
}

// Método que transfoma um Slice de entidades em Slice de models PartReq
func (*PartReq) TransformToSliceOfModel(e []entities.PartsReq) *[]PartReq {

	var (
		m  []PartReq
		wg sync.WaitGroup
	)

	channel := make(chan int)

	channel <- 0

	for range e {

		wg.Add(1)

		go func() {

			m[<-channel].ID = e[<-channel].Id
			m[<-channel].ProductId = e[<-channel].ProductId
			m[<-channel].ProductName = e[<-channel].ProductName
			m[<-channel].ProductPrice = e[<-channel].ProductPrice
			m[<-channel].ProdtQtyReq = e[<-channel].ProdtQtyReq
			m[<-channel].SaleID = e[<-channel].SaleId

			channel <- <-channel + 1

			wg.Done()
		}()

		if index := <-channel + 1; len(m)-1 > index {

			wg.Add(1)

			go func() {
				m[<-channel+1].ID = e[<-channel+1].Id
				m[<-channel+1].ProductId = e[<-channel+1].ProductId
				m[<-channel+1].ProductName = e[<-channel+1].ProductName
				m[<-channel+1].ProductPrice = e[<-channel+1].ProductPrice
				m[<-channel+1].ProdtQtyReq = e[<-channel+1].ProdtQtyReq
				m[<-channel+1].SaleID = e[<-channel+1].SaleId

				channel <- <-channel + 1

				wg.Done()
			}()
		}

		wg.Wait()
	}

	close(channel)

	return &m

}
