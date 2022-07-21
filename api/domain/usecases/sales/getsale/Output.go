package getsale

import (
	"sync"

	"github.com/realnfcs/ultividros-project/api/domain/entities"
)

// Usecase Output Port responsável pelos dados que serão retornados
type Output struct {
	OutputData `json:"data"`
	Status     int    `json:"-"`
	Err        string `json:"error"`
}

// OutpurData responsável pelo dado da entity em si que será passado pelas camadas externas
type OutputData struct {
	Id            string          `json:"id"`
	ClientId      string          `json:"client_id"`
	CommonGlssReq []CommonGlssReq `json:"common_glss_req"`
	PartsReq      []PartsReq      `json:"parts_req"`
	TempGlssReq   []TempGlssReq   `json:"temp_glss"`
	IsActive      bool            `json:"is_active"`
}

// Informações contidas dentro do OutputData
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

type TempGlssReq struct {
	ProductsRequest
}

type PartsReq struct {
	ProductsRequest
}

type CommonGlssReq struct {
	ProductsRequest
	RequestWidth  float32 `json:"request_width"`
	RequestHeight float32 `json:"request_height"`
}

func (*Output) Init(e *entities.Sale, status int, err error) *Output {
	if e != nil {

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

		return &Output{
			OutputData: OutputData{
				Id:            e.Id,
				ClientId:      e.ClientId,
				CommonGlssReq: comnGlssReq,
				PartsReq:      partReq,
				TempGlssReq:   tempGlssReq,
				IsActive:      e.IsActive,
			},
			Status: status,
			Err:    "",
		}
	}

	return &Output{OutputData{}, status, err.Error()}
}
