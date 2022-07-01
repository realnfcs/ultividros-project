package savesale

import (
	"sync"

	"github.com/realnfcs/ultividros-project/api/domain/entities"
)

// Usecase Input Port responsável pelos dados que entrarão
type Input struct {
	ClientId   string            `json:"client_id"`
	Products   []ProductsRequest `json:"products_request"`
	CommonGlss []CommonGlssReq   `json:"common_glss"`
}

// Produtos requisitados pelo cliente que virão no input
type ProductsRequest struct {
	Id           string  `json:"id"`
	ProductId    string  `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float32 `json:"product_price"`
	ProdtQtyReq  uint32  `json:"prodt_qty_req"`
	SaleId       string  `json:"sale_id"`
}

// Nenhum ou muitos vidros comuns requisitados pelo cliente que virão no input
type CommonGlssReq struct {
	ProductsRequest
	RequestWidth  float32 `json:"request_width"`
	RequestHeight float32 `json:"request_height"`
}

// Método que converte um input na entidade Sale
func (i *Input) ConvertToSale() *entities.Sale {
	prods := make([]string, (len(i.Products) + len(i.CommonGlss)))

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		if len(i.Products) != 0 {
			for i, v := range i.Products {
				prods[i] = v.ProductId
			}
		}
		wg.Done()
	}()

	if len(i.CommonGlss) != 0 {
		for i, v := range i.CommonGlss {
			prods[i] = v.ProductId
		}
	}

	wg.Wait()

	return &entities.Sale{
		"",
		i.ClientId,
		prods,
	}
}

// Método que converte os produtos do input na entidade ProductsRequest
func (i *Input) ConvertProdInputInProdEnt() *[]entities.ProductsRequest {

	if len(i.Products) == 0 {
		return nil
	}

	prods := make([]entities.ProductsRequest, len(i.Products))

	if len(i.Products) > 5 {

		var wg sync.WaitGroup
		channel := make(chan int)

		for i, v := range i.Products {

			channel <- i

			wg.Add(1)

			go func(ch chan int, p ProductsRequest) {
				prods[<-channel].Id = p.Id
				prods[<-channel].ProductId = p.ProductId
				prods[<-channel].ProductName = p.ProductName
				prods[<-channel].ProductPrice = p.ProductPrice
				prods[<-channel].ProdtQtyReq = p.ProdtQtyReq
				prods[<-channel].SaleId = p.SaleId

				wg.Done()
				channel <- <-channel + 1
			}(channel, v)
		}

		wg.Wait()

		close(channel)

		return &prods
	}

	for i, v := range i.Products {
		prods[i].Id = v.Id
		prods[i].ProductId = v.ProductId
		prods[i].ProductName = v.ProductName
		prods[i].ProductPrice = v.ProductPrice
		prods[i].ProdtQtyReq = v.ProdtQtyReq
		prods[i].SaleId = v.SaleId
	}

	return &prods
}

// Método que converte os vidros comuns na entidade ComnGlssReq
func (i *Input) ConvertComnGlssReqInputInComnGlssReqEnt() *[]entities.CommonGlssReq {

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
				comnGlss[<-channel].ProdtQtyReq = c.ProdtQtyReq
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
		comnGlss[i].ProdtQtyReq = v.ProdtQtyReq
		comnGlss[i].RequestWidth = v.RequestWidth
		comnGlss[i].RequestHeight = v.RequestHeight
		comnGlss[i].SaleId = v.SaleId
	}

	return &comnGlss
}
