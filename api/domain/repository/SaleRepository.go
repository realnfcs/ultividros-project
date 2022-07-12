package repository

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// A interface SaleRepository provém os métodos necessários para
// que o repositório a ser usado realiza o objetivo do usecase
type SaleRepository interface {
	GetSales() (*[]entities.Sale, int, error)
	SaveSale(entities.Sale /*, []entities.PartsReq, []entities.CommonGlssReq, []entities.TempGlssReq */) (string, int, error)
}
