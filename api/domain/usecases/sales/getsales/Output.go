package getsales

import (
	"sync"

	"github.com/realnfcs/ultividros-project/api/domain/entities"
)

// Usecase Output Port responsável pelos dados que serão retornados
type Output struct {
	Data   []OutputData `json:"data"`
	Status int          `json:"-"`
	Err    string       `json:"error"`
}

// OutputData é a estrutura de dados que será retornado em um array
// no Output
type OutputData struct {
	Id            string          `json:"id"`
	ClientId      string          `json:"client_id"`
	CommonGlssReq []CommonGlssReq `json:"common_glss_req"`
	PartsReq      []PartsReq      `json:"parts_req"`
	TempGlssReq   []TempGlssReq   `json:"temp_glss"`
	IsActive      bool            `json:"is_active"`
}

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

func (*Output) Init(i *[]entities.Sale, status int, err error) *Output {

	if i == nil || err != nil {
		return &Output{[]OutputData{}, status, err.Error()}
	}

	output := make([]OutputData, len(*i))

	var wg sync.WaitGroup

	for i, v := range *i {
		output[i].Id = v.Id
		output[i].ClientId = v.ClientId
		output[i].IsActive = v.IsActive

		comnGlss := make([]CommonGlssReq, len(v.CommonGlssReq))

		if len(v.CommonGlssReq) > 0 {
			wg.Add(1)

			go func() {
				for index, value := range v.CommonGlssReq {
					if value.SaleId == v.Id {
						comnGlss[index].Id = value.Id
						comnGlss[index].ProductId = value.ProductId
						comnGlss[index].ProductName = value.ProductName
						comnGlss[index].ProductPrice = value.ProductPrice
						comnGlss[index].ProdtQtyReq = value.ProdtQtyReq
						comnGlss[index].WasCancelled = value.WasCancelled
						comnGlss[index].WasConfirmed = value.WasConfirmed
						comnGlss[index].RequestWidth = value.RequestWidth
						comnGlss[index].RequestHeight = value.RequestHeight
						comnGlss[index].SaleId = value.SaleId
					}
				}
				wg.Done()
			}()
		}

		partsReq := make([]PartsReq, len(v.PartsReq))

		if len(v.PartsReq) > 0 {
			wg.Add(1)

			go func() {
				for index, value := range v.PartsReq {
					if value.SaleId == v.Id {
						partsReq[index].Id = value.Id
						partsReq[index].ProductId = value.ProductId
						partsReq[index].ProductName = value.ProductName
						partsReq[index].ProductPrice = value.ProductPrice
						partsReq[index].ProdtQtyReq = value.ProdtQtyReq
						partsReq[index].WasCancelled = value.WasCancelled
						partsReq[index].WasConfirmed = value.WasConfirmed
						partsReq[index].SaleId = value.SaleId
					}
				}
				wg.Done()
			}()
		}

		tempGlss := make([]TempGlssReq, len(v.TempGlssReq))

		if len(v.TempGlssReq) > 0 {
			wg.Add(1)

			go func() {
				for index, value := range v.TempGlssReq {
					if value.SaleId == v.Id {
						tempGlss[index].Id = value.Id
						tempGlss[index].ProductId = value.ProductId
						tempGlss[index].ProductName = value.ProductName
						tempGlss[index].ProductPrice = value.ProductPrice
						tempGlss[index].ProdtQtyReq = value.ProdtQtyReq
						tempGlss[index].WasCancelled = value.WasCancelled
						tempGlss[index].WasConfirmed = value.WasConfirmed
						tempGlss[index].SaleId = value.SaleId
					}
				}
				wg.Done()
			}()
		}

		wg.Wait()

		output[i].CommonGlssReq = comnGlss
		output[i].PartsReq = partsReq
		output[i].TempGlssReq = tempGlss
	}

	return &Output{output, status, ""}
}
