package repository

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// A interface SaleRepository provém os métodos necessários para
// que o repositório a ser usado realiza o objetivo do usecase
type SaleRepository interface {
	SaveSale(entities.Sale, []entities.ProductsRequest, []entities.CommonGlssReq) (string, int, error)
}
