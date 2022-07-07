package controllers

import (
	"github.com/realnfcs/ultividros-project/api/domain/repository"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/sales/savesale"
)

// Struct para iniciar a controller das vendas necessitando
// de um repository para funcionar
type SaleController struct {
	Repo                   repository.SaleRepository
	CommonGlssRepository   repository.CommonGlassRepository
	PartRepository         repository.PartRepository
	TemperedGlssRepository repository.TemperedGlassRepository
}

// Método da controller que comunica com o usecase para salvar um objeto de venda
// de acordo com os dados passados no parâmetro
func (s *SaleController) SaveSale(i savesale.Input) *savesale.Output {
	saveSale := savesale.SaveSale{SaleRepository: s.Repo, CommonGlssRepository: s.CommonGlssRepository, PartRepository: s.PartRepository, TemperedGlssRepository: s.TemperedGlssRepository}
	output := saveSale.Execute(i)
	return output
}
