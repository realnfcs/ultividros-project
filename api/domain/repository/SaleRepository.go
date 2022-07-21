package repository

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// A interface SaleRepository provém os métodos necessários para
// que o repositório a ser usado realiza o objetivo do usecase
type SaleRepository interface {
	GetSale(string) (*entities.Sale, int, error)
	GetSales() (*[]entities.Sale, int, error)
	SaveSale(entities.Sale) (string, int, error)
	PatchSale(entities.Sale) (string, int, error)
	CloseSale(entities.Sale) (string, int, error)
}
