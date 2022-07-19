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

// Essa struct provém as informações mais específicas dos vidros comuns
// requisitados na venda porém com mais enfâse nas bibliotecas usadas.
type CommonGlssReq struct {
	ProductsRequest
	RequestWidth  float32 `json:"request_width"`
	RequestHeight float32 `json:"request_height"`
}

// Método para criar um uuid antes de salvar no banco de dados
func (m *CommonGlssReq) BeforeCreate(scope *gorm.DB) error {
	id := uuid.New().String()
	if id == "" {
		return errors.New("Cannot create uuid")
	}

	m.ID = strings.Replace(id, "-", "", -1)

	return nil
}

// Método responsável por transformar um Model em uma entidades CommonGlssReq
func (m *CommonGlssReq) TranformToEntity() *entities.CommonGlssReq {
	return &entities.CommonGlssReq{
		ProductsRequest: entities.ProductsRequest{
			Id:           m.ID,
			ProductId:    m.ProductId,
			ProductName:  m.ProductName,
			ProductPrice: m.ProductPrice,
			ProdtQtyReq:  m.ProdtQtyReq,
			WasCancelled: m.WasCancelled,
			WasConfirmed: m.WasConfirmed,
			SaleId:       m.SaleID,
		},
		RequestWidth:  m.RequestWidth,
		RequestHeight: m.RequestHeight,
	}
}

// Método responsável por transformar um Slice de Models em um Slice de entidades CommonGlssReq
func (*CommonGlssReq) TransformToSliceOfEntity(m []CommonGlssReq) *[]entities.CommonGlssReq {

	comnGlss := make([]entities.CommonGlssReq, len(m))

	var wg sync.WaitGroup

	channel := make(chan int)

	channel <- 0

	for range m {

		wg.Add(1)

		go func() {

			comnGlss[<-channel].Id = m[<-channel].ID
			comnGlss[<-channel].ProductId = m[<-channel].ProductId
			comnGlss[<-channel].ProductName = m[<-channel].ProductName
			comnGlss[<-channel].ProductPrice = m[<-channel].ProductPrice
			comnGlss[<-channel].ProdtQtyReq = m[<-channel].ProdtQtyReq
			comnGlss[<-channel].WasCancelled = m[<-channel].WasCancelled
			comnGlss[<-channel].WasConfirmed = m[<-channel].WasConfirmed
			comnGlss[<-channel].SaleId = m[<-channel].SaleID
			comnGlss[<-channel].RequestWidth = m[<-channel].RequestWidth
			comnGlss[<-channel].RequestHeight = m[<-channel].RequestHeight

			channel <- <-channel + 1

			wg.Done()
		}()

		if index := <-channel + 1; len(m)-1 > index {

			wg.Add(1)

			go func() {
				comnGlss[<-channel+1].Id = m[<-channel+1].ID
				comnGlss[<-channel+1].ProductId = m[<-channel+1].ProductId
				comnGlss[<-channel+1].ProductName = m[<-channel+1].ProductName
				comnGlss[<-channel+1].ProductPrice = m[<-channel+1].ProductPrice
				comnGlss[<-channel+1].ProdtQtyReq = m[<-channel+1].ProdtQtyReq
				comnGlss[<-channel+1].WasCancelled = m[<-channel+1].WasCancelled
				comnGlss[<-channel+1].WasConfirmed = m[<-channel+1].WasConfirmed
				comnGlss[<-channel+1].SaleId = m[<-channel+1].SaleID
				comnGlss[<-channel+1].RequestWidth = m[<-channel+1].RequestWidth
				comnGlss[<-channel+1].RequestHeight = m[<-channel+1].RequestHeight

				channel <- <-channel + 1

				wg.Done()
			}()
		}

		wg.Wait()
	}

	close(channel)

	return &comnGlss
}

// Método responsável por transformar a entidade ComnGlssReq em model
func (m *CommonGlssReq) TransformToModel(e entities.CommonGlssReq) *CommonGlssReq {
	return &CommonGlssReq{
		ProductsRequest{
			e.Id,
			e.ProductId,
			e.ProductName,
			e.ProductPrice,
			e.ProdtQtyReq,
			e.WasCancelled,
			e.WasConfirmed,
			e.SaleId,
			time.Time{},
			time.Time{},
			sql.NullTime{},
		},
		e.RequestWidth,
		e.RequestHeight,
	}
}

// Método que transfoma um Slice de entidades em Slice de models CommonGlssReq
func (*CommonGlssReq) TransformToSliceOfModel(e []entities.CommonGlssReq) *[]CommonGlssReq {

	var (
		m  []CommonGlssReq
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
			m[<-channel].WasCancelled = e[<-channel].WasCancelled
			m[<-channel].WasConfirmed = e[<-channel].WasConfirmed
			m[<-channel].SaleID = e[<-channel].SaleId
			m[<-channel].RequestWidth = e[<-channel].RequestWidth
			m[<-channel].RequestHeight = e[<-channel].RequestHeight

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
				m[<-channel+1].WasCancelled = e[<-channel+1].WasCancelled
				m[<-channel+1].WasConfirmed = e[<-channel+1].WasConfirmed
				m[<-channel+1].SaleID = e[<-channel+1].SaleId
				m[<-channel+1].RequestWidth = e[<-channel+1].RequestWidth
				m[<-channel+1].RequestHeight = e[<-channel+1].RequestHeight

				channel <- <-channel + 1

				wg.Done()
			}()
		}

		wg.Wait()
	}

	close(channel)

	return &m

}
