package patchsale

import (
	"sync"

	"github.com/realnfcs/ultividros-project/api/domain/entities"
)

// Usecase Input Port responsável pelos dados que entrarão
type Input struct {
	Id            string          `json:"id"`
	ClientId      string          `json:"client_id"`
	CommonGlssReq []CommonGlssReq `json:"common_glss_req"`
	PartsReq      []PartsReq      `json:"parts_req"`
	TempGlssReq   []TempGlssReq   `json:"temp_glss"`
	IsActive      bool            `json:"is_active"`
}

// Produtos requisitados pelo cliente que virão no input
type ProductsRequest struct {
	Id           string  `json:"id"`
	ProductId    string  `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float32 `json:"product_price"`
	ProdtQtyReq  uint32  `json:"prodt_qty_req"`
	WasCancelled bool    `json:"was_cancelled"`
	WasConfirmed bool    `json:"was_confimed"`
	SaleId       string  `json:"sale_id"`
}

// Nenhum ou muitos vidros temperados requisitados pelo cliente que virão no input
type TempGlssReq struct {
	ProductsRequest
}

// Nenhum ou muitos peças requisitados pelo cliente que virão no input
type PartsReq struct {
	ProductsRequest
}

// Nenhum ou muitos vidros comuns requisitados pelo cliente que virão no input
type CommonGlssReq struct {
	ProductsRequest
	RequestWidth  float32 `json:"request_width"`
	RequestHeight float32 `json:"request_height"`
}

func (*Input) Init(e entities.Sale, userId string) *Input {

	var wg sync.WaitGroup

	comnGlssReq := make([]CommonGlssReq, len(e.CommonGlssReq))

	if len(e.CommonGlssReq) > 0 {
		wg.Add(1)

		go func() {
			for i, v := range e.CommonGlssReq {
				comnGlssReq[i].Id = v.Id
				comnGlssReq[i].ProductId = v.ProductId
				comnGlssReq[i].ProductName = v.ProductName
				comnGlssReq[i].ProductPrice = v.ProductPrice
				comnGlssReq[i].ProdtQtyReq = v.ProdtQtyReq
				comnGlssReq[i].WasCancelled = v.WasCancelled
				comnGlssReq[i].WasConfirmed = v.WasConfirmed
				comnGlssReq[i].RequestWidth = v.RequestWidth
				comnGlssReq[i].RequestHeight = v.RequestHeight
				comnGlssReq[i].SaleId = v.SaleId
			}

			wg.Done()
		}()
	}

	partReq := make([]PartsReq, len(e.PartsReq))

	if len(e.PartsReq) > 0 {
		wg.Add(1)

		go func() {
			for i, v := range e.PartsReq {
				partReq[i].Id = v.Id
				partReq[i].ProductId = v.ProductId
				partReq[i].ProductName = v.ProductName
				partReq[i].ProductPrice = v.ProductPrice
				partReq[i].ProdtQtyReq = v.ProdtQtyReq
				partReq[i].WasCancelled = v.WasCancelled
				partReq[i].WasConfirmed = v.WasConfirmed
				partReq[i].SaleId = v.SaleId
			}

			wg.Done()
		}()
	}

	tempGlssReq := make([]TempGlssReq, len(e.TempGlssReq))

	if len(e.TempGlssReq) > 0 {
		wg.Add(1)

		go func() {
			for i, v := range e.TempGlssReq {
				tempGlssReq[i].Id = v.Id
				tempGlssReq[i].ProductId = v.ProductId
				tempGlssReq[i].ProductName = v.ProductName
				tempGlssReq[i].ProductPrice = v.ProductPrice
				tempGlssReq[i].ProdtQtyReq = v.ProdtQtyReq
				tempGlssReq[i].WasCancelled = v.WasCancelled
				tempGlssReq[i].WasConfirmed = v.WasConfirmed
				tempGlssReq[i].SaleId = v.SaleId
			}

			wg.Done()
		}()
	}

	return &Input{
		e.Id,
		e.ClientId,
		comnGlssReq,
		partReq,
		tempGlssReq,
		e.IsActive,
	}
}

// Método responsável em converter um input em uma entidade de venda
func (i *Input) ConvertToSale() *entities.Sale {

	var wg sync.WaitGroup

	comnGlssReq := make([]entities.CommonGlssReq, len(i.CommonGlssReq))

	if len(i.CommonGlssReq) > 0 {
		wg.Add(1)

		go func() {
			for i, v := range i.CommonGlssReq {
				comnGlssReq[i].Id = v.Id
				comnGlssReq[i].ProductId = v.ProductId
				comnGlssReq[i].ProductName = v.ProductName
				comnGlssReq[i].ProductPrice = v.ProductPrice
				comnGlssReq[i].ProdtQtyReq = v.ProdtQtyReq
				comnGlssReq[i].WasCancelled = v.WasCancelled
				comnGlssReq[i].WasConfirmed = v.WasConfirmed
				comnGlssReq[i].RequestWidth = v.RequestWidth
				comnGlssReq[i].RequestHeight = v.RequestHeight
				comnGlssReq[i].SaleId = v.SaleId
			}

			wg.Done()
		}()
	}

	partReq := make([]entities.PartsReq, len(i.PartsReq))

	if len(i.PartsReq) > 0 {
		wg.Add(1)

		go func() {
			for i, v := range i.PartsReq {
				partReq[i].Id = v.Id
				partReq[i].ProductId = v.ProductId
				partReq[i].ProductName = v.ProductName
				partReq[i].ProductPrice = v.ProductPrice
				partReq[i].ProdtQtyReq = v.ProdtQtyReq
				partReq[i].WasCancelled = v.WasCancelled
				partReq[i].WasConfirmed = v.WasConfirmed
				partReq[i].SaleId = v.SaleId
			}

			wg.Done()
		}()
	}

	tempGlssReq := make([]entities.TempGlssReq, len(i.TempGlssReq))

	if len(i.TempGlssReq) > 0 {
		wg.Add(1)

		go func() {
			for i, v := range i.TempGlssReq {
				tempGlssReq[i].Id = v.Id
				tempGlssReq[i].ProductId = v.ProductId
				tempGlssReq[i].ProductName = v.ProductName
				tempGlssReq[i].ProductPrice = v.ProductPrice
				tempGlssReq[i].ProdtQtyReq = v.ProdtQtyReq
				tempGlssReq[i].WasCancelled = v.WasCancelled
				tempGlssReq[i].WasConfirmed = v.WasConfirmed
				tempGlssReq[i].SaleId = v.SaleId
			}

			wg.Done()
		}()
	}

	wg.Wait()

	return &entities.Sale{
		Id:            i.Id,
		ClientId:      i.ClientId,
		CommonGlssReq: comnGlssReq,
		PartsReq:      partReq,
		TempGlssReq:   tempGlssReq,
		IsActive:      i.IsActive,
	}
}
