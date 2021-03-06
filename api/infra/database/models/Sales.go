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
	ID            string          `json:"id" gorm:"primaryKey"`
	ClientID      string          `json:"client_id"`
	CommonGlssReq []CommonGlssReq `json:"common_glss_req"`
	PartReq       []PartReq       `json:"part_req"`
	TempGlssReq   []TempGlssReq   `json:"temp_glss_req"`
	IsActive      bool            `json:"is_active" gorm:"default:true"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	DeletedAt     sql.NullTime    `json:"deleted_at" gorm:"index"`
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

	comnGlss := make([]entities.CommonGlssReq, len(m.CommonGlssReq))
	partReq := make([]entities.PartsReq, len(m.PartReq))
	tempGlss := make([]entities.TempGlssReq, len(m.TempGlssReq))

	var wg sync.WaitGroup

	if len(m.CommonGlssReq) > 0 {
		wg.Add(1)
		go func() {
			for i, v := range m.CommonGlssReq {
				comnGlss[i].Id = v.ID
				comnGlss[i].ProductId = v.ProductId
				comnGlss[i].ProductName = v.ProductName
				comnGlss[i].ProductPrice = v.ProductPrice
				comnGlss[i].ProdtQtyReq = v.ProdtQtyReq
				comnGlss[i].ProdtQtyReq = v.ProdtQtyReq
				comnGlss[i].WasCancelled = v.WasCancelled
				comnGlss[i].WasConfirmed = v.WasConfirmed
				comnGlss[i].RequestWidth = v.RequestWidth
				comnGlss[i].RequestHeight = v.RequestHeight
				comnGlss[i].SaleId = v.SaleID
			}
			wg.Done()
		}()
	}

	if len(m.PartReq) > 0 {
		wg.Add(1)
		go func() {
			for i, v := range m.PartReq {
				partReq[i].Id = v.ID
				partReq[i].ProductId = v.ProductId
				partReq[i].ProductName = v.ProductName
				partReq[i].ProductPrice = v.ProductPrice
				partReq[i].ProdtQtyReq = v.ProdtQtyReq
				partReq[i].WasCancelled = v.WasCancelled
				partReq[i].WasConfirmed = v.WasConfirmed
				partReq[i].SaleId = v.SaleID
			}
			wg.Done()
		}()
	}

	if len(m.TempGlssReq) > 0 {
		wg.Add(1)
		go func() {
			for i, v := range m.TempGlssReq {
				tempGlss[i].Id = v.ID
				tempGlss[i].ProductId = v.ProductId
				tempGlss[i].ProductName = v.ProductName
				tempGlss[i].ProductPrice = v.ProductPrice
				tempGlss[i].ProdtQtyReq = v.ProdtQtyReq
				tempGlss[i].WasCancelled = v.WasCancelled
				tempGlss[i].WasConfirmed = v.WasConfirmed
				tempGlss[i].SaleId = v.SaleID
			}
			wg.Done()
		}()
	}

	wg.Wait()

	return &entities.Sale{
		Id:            m.ID,
		ClientId:      m.ClientID,
		CommonGlssReq: comnGlss,
		PartsReq:      partReq,
		TempGlssReq:   tempGlss,
		IsActive:      m.IsActive,
	}
}

// Método responsável por transformar um Slice de Models em um Slice de entidades Sale
func (*Sale) TransformToSliceOfEntity(m []Sale) *[]entities.Sale {

	sale := make([]entities.Sale, len(m))

	for i, v := range m {
		sale[i] = *v.TranformToEntity()
	}

	return &sale

}

// Método responsável por transformar a entidade Sale em model
func (m *Sale) TransformToModel(e entities.Sale) *Sale {

	comnGlss := make([]CommonGlssReq, len(e.CommonGlssReq))
	parts := make([]PartReq, len(e.PartsReq))
	tempGlss := make([]TempGlssReq, len(e.TempGlssReq))

	if len(e.CommonGlssReq) > 0 {
		for i, v := range e.CommonGlssReq {
			comnGlss[i].ID = v.Id
			comnGlss[i].ProductId = v.ProductId
			comnGlss[i].ProductName = v.ProductName
			comnGlss[i].ProductPrice = v.ProductPrice
			comnGlss[i].ProdtQtyReq = v.ProdtQtyReq
			comnGlss[i].WasCancelled = v.WasCancelled
			comnGlss[i].WasConfirmed = v.WasConfirmed
			comnGlss[i].RequestWidth = v.RequestWidth
			comnGlss[i].RequestHeight = v.RequestHeight
			comnGlss[i].SaleID = v.SaleId
		}
	}

	if len(e.PartsReq) > 0 {
		for i, v := range e.PartsReq {
			parts[i].ID = v.Id
			parts[i].ProductId = v.ProductId
			parts[i].ProductName = v.ProductName
			parts[i].ProductPrice = v.ProductPrice
			parts[i].ProdtQtyReq = v.ProdtQtyReq
			parts[i].WasCancelled = v.WasCancelled
			parts[i].WasConfirmed = v.WasConfirmed
			parts[i].SaleID = v.SaleId
		}
	}

	if len(e.TempGlssReq) > 0 {
		for i, v := range e.TempGlssReq {
			tempGlss[i].ID = v.Id
			tempGlss[i].ProductId = v.ProductId
			tempGlss[i].ProductName = v.ProductName
			tempGlss[i].ProductPrice = v.ProductPrice
			tempGlss[i].ProdtQtyReq = v.ProdtQtyReq
			tempGlss[i].WasCancelled = v.WasCancelled
			tempGlss[i].WasConfirmed = v.WasConfirmed
			tempGlss[i].SaleID = v.SaleId
		}
	}

	return &Sale{
		e.Id,
		e.ClientId,
		comnGlss,
		parts,
		tempGlss,
		e.IsActive,
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

			m[<-channel].TransformToModel(e[<-channel])

			channel <- <-channel + 1

			wg.Done()
		}()

		if index := <-channel + 1; len(e)-1 > index {

			wg.Add(1)

			go func() {

				m[<-channel+1].TransformToModel(e[<-channel+1])

				channel <- <-channel + 1

				wg.Done()
			}()
		}

		wg.Wait()
	}

	close(channel)

	return &m
}
