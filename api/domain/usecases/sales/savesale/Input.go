package savesale

import (
	"sync"

	"github.com/realnfcs/ultividros-project/api/domain/entities"
)

// Usecase Input Port responsável pelos dados que entrarão
type Input struct {
	ClientId   string          `json:"client_id"`
	Parts      []PartsReq      `json:"parts_req"`
	TempGlss   []TempGlssReq   `json:"temp_glss_req"`
	CommonGlss []CommonGlssReq `json:"common_glss_req"`
	IsActive   bool            `json:"is_active"`
}

// Produtos requisitados pelo cliente que virão no input
type ProductsRequest struct {
	Id           string  `json:"id"`
	ProductId    string  `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float32 `json:"product_price"`
	ProdQtyReq   uint32  `json:"prod_qty_req"`
	WasCancelled bool    `json:"was_cancelled"`
	WasConfirmed bool    `json:"was_confimed"`
	SaleId       string  `json:"sale_id"`
}

// Nenhum ou muitos peças requisitados pelo cliente que virão no input
type PartsReq struct {
	ProductsRequest
}

// Nenhum ou muitos vidros temperados requisitados pelo cliente que virão no input
type TempGlssReq struct {
	ProductsRequest
}

// Nenhum ou muitos vidros comuns requisitados pelo cliente que virão no input
type CommonGlssReq struct {
	ProductsRequest
	RequestWidth  float32 `json:"request_width"`
	RequestHeight float32 `json:"request_height"`
}

// Método que converte um input na entidade Sale
func (i *Input) ConvertToSale() *entities.Sale {

	comnGlss := make([]entities.CommonGlssReq, len(i.CommonGlss))
	parts := make([]entities.PartsReq, len(i.Parts))
	tempGlss := make([]entities.TempGlssReq, len(i.TempGlss))

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		if len(i.Parts) != 0 {
			for index, v := range i.Parts {
				parts[index].Id = v.Id
				parts[index].ProductId = v.ProductId
				parts[index].ProductName = v.ProductName
				parts[index].ProductPrice = v.ProductPrice
				parts[index].ProdtQtyReq = v.ProdQtyReq
				parts[index].WasCancelled = v.WasCancelled
				parts[index].WasConfirmed = v.WasConfirmed
				parts[index].SaleId = v.SaleId
			}
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		if len(i.TempGlss) != 0 {
			for index, v := range i.TempGlss {
				tempGlss[index].Id = v.Id
				tempGlss[index].ProductId = v.ProductId
				tempGlss[index].ProductName = v.ProductName
				tempGlss[index].ProductPrice = v.ProductPrice
				tempGlss[index].ProdtQtyReq = v.ProdQtyReq
				tempGlss[index].WasCancelled = v.WasCancelled
				tempGlss[index].WasConfirmed = v.WasConfirmed
				tempGlss[index].SaleId = v.SaleId
			}
		}
		wg.Done()
	}()

	if len(i.CommonGlss) != 0 {
		for index, v := range i.CommonGlss {
			comnGlss[index].Id = v.Id
			comnGlss[index].ProductId = v.ProductId
			comnGlss[index].ProductName = v.ProductName
			comnGlss[index].ProductPrice = v.ProductPrice
			comnGlss[index].ProdtQtyReq = v.ProdQtyReq
			comnGlss[index].WasCancelled = v.WasCancelled
			comnGlss[index].WasConfirmed = v.WasConfirmed
			comnGlss[index].RequestWidth = v.RequestWidth
			comnGlss[index].RequestHeight = v.RequestHeight
			comnGlss[index].SaleId = v.SaleId
		}
	}

	wg.Wait()

	return &entities.Sale{
		Id:            "",
		ClientId:      i.ClientId,
		CommonGlssReq: comnGlss,
		PartsReq:      parts,
		TempGlssReq:   tempGlss,
		IsActive:      i.IsActive,
	}
}

// Método que converte os produtos do input na entidade de vidros temperados (TempGlssReq)
func (i *Input) ConvertTempGlssReqInputInEnt() *[]entities.TempGlssReq {

	if len(i.TempGlss) == 0 {
		return nil
	}

	prods := make([]entities.TempGlssReq, len(i.TempGlss))

	if len(i.TempGlss) > 5 {

		var wg sync.WaitGroup
		channel := make(chan int)

		channel <- 0

		for _, v := range i.TempGlss {

			wg.Add(1)

			go func(ch chan int, p TempGlssReq) {
				prods[<-channel].Id = p.Id
				prods[<-channel].ProductId = p.ProductId
				prods[<-channel].ProductName = p.ProductName
				prods[<-channel].ProductPrice = p.ProductPrice
				prods[<-channel].ProdtQtyReq = p.ProdQtyReq
				prods[<-channel].WasCancelled = p.WasCancelled
				prods[<-channel].WasConfirmed = p.WasConfirmed
				prods[<-channel].SaleId = p.SaleId

				wg.Done()
				channel <- <-channel + 1
			}(channel, v)
		}

		wg.Wait()

		close(channel)

		return &prods
	}

	for i, v := range i.TempGlss {
		prods[i].Id = v.Id
		prods[i].ProductId = v.ProductId
		prods[i].ProductName = v.ProductName
		prods[i].ProductPrice = v.ProductPrice
		prods[i].ProdtQtyReq = v.ProdQtyReq
		prods[i].WasCancelled = v.WasCancelled
		prods[i].WasConfirmed = v.WasConfirmed
		prods[i].SaleId = v.SaleId
	}

	return &prods
}

// Método que converte os vidros comuns na entidade de vidros comuns (ComnGlssReq)
func (i *Input) ConvertComnGlssReqInputInEnt() *[]entities.CommonGlssReq {

	if len(i.CommonGlss) == 0 {
		return nil
	}

	comnGlss := make([]entities.CommonGlssReq, len(i.CommonGlss))

	if len(i.CommonGlss) > 5 {

		var wg sync.WaitGroup
		channel := make(chan int)

		for i, v := range i.CommonGlss {

			channel <- i

			wg.Add(1)

			go func(ch chan int, c CommonGlssReq) {
				comnGlss[<-channel].Id = c.Id
				comnGlss[<-channel].ProductId = c.ProductId
				comnGlss[<-channel].ProductName = c.ProductName
				comnGlss[<-channel].ProductPrice = c.ProductPrice
				comnGlss[<-channel].ProdtQtyReq = c.ProdQtyReq
				comnGlss[<-channel].WasCancelled = c.WasCancelled
				comnGlss[<-channel].WasConfirmed = c.WasConfirmed
				comnGlss[<-channel].RequestWidth = c.RequestWidth
				comnGlss[<-channel].RequestHeight = c.RequestHeight
				comnGlss[<-channel].SaleId = c.SaleId

				wg.Done()

				channel <- <-channel + 1
			}(channel, v)
		}

		wg.Wait()

		close(channel)

		return &comnGlss
	}

	for i, v := range i.CommonGlss {
		comnGlss[i].Id = v.Id
		comnGlss[i].ProductId = v.ProductId
		comnGlss[i].ProductName = v.ProductName
		comnGlss[i].ProductPrice = v.ProductPrice
		comnGlss[i].ProdtQtyReq = v.ProdQtyReq
		comnGlss[i].WasCancelled = v.WasCancelled
		comnGlss[i].WasConfirmed = v.WasConfirmed
		comnGlss[i].RequestWidth = v.RequestWidth
		comnGlss[i].RequestHeight = v.RequestHeight
		comnGlss[i].SaleId = v.SaleId
	}

	return &comnGlss
}

// Método que converte os produtos do input na entidade de peças (PartsReq)
func (i *Input) ConvertPartReqInputInEnt() *[]entities.PartsReq {

	if len(i.Parts) == 0 {
		return nil
	}

	prods := make([]entities.PartsReq, len(i.Parts))

	if len(i.Parts) > 5 {

		var wg sync.WaitGroup
		channel := make(chan int)

		channel <- 0

		for _, v := range i.Parts {

			wg.Add(1)

			go func(ch chan int, p PartsReq) {
				prods[<-channel].Id = p.Id
				prods[<-channel].ProductId = p.ProductId
				prods[<-channel].ProductName = p.ProductName
				prods[<-channel].ProductPrice = p.ProductPrice
				prods[<-channel].ProdtQtyReq = p.ProdQtyReq
				prods[<-channel].WasCancelled = p.WasCancelled
				prods[<-channel].WasConfirmed = p.WasConfirmed
				prods[<-channel].SaleId = p.SaleId

				wg.Done()
				channel <- <-channel + 1
			}(channel, v)
		}

		wg.Wait()

		close(channel)

		return &prods
	}

	for i, v := range i.Parts {
		prods[i].Id = v.Id
		prods[i].ProductId = v.ProductId
		prods[i].ProductName = v.ProductName
		prods[i].ProductPrice = v.ProductPrice
		prods[i].ProdtQtyReq = v.ProdQtyReq
		prods[i].WasCancelled = v.WasCancelled
		prods[i].WasConfirmed = v.WasConfirmed
		prods[i].SaleId = v.SaleId
	}

	return &prods
}
