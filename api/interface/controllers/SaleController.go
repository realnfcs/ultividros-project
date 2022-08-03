package controllers

import (
	"github.com/realnfcs/ultividros-project/api/domain/repository"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/sales/closesale"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/sales/deletesale"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/sales/getsale"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/sales/getsales"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/sales/patchsale"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/sales/savesale"
)

// Struct para iniciar a controller das vendas necessitando
// de um repository para funcionar
type SaleController struct {
	Repo                   repository.SaleRepository
	CommonGlssRepository   repository.CommonGlassRepository
	PartRepository         repository.PartRepository
	TemperedGlssRepository repository.TemperedGlassRepository
	UserRepository         repository.UserRepository
}

// Método da controller que comunica com o usecase para a obtenção de dados de uma
// única venda e seus produtos trazendo a resposta ao cliente
func (s *SaleController) GetSale(i getsale.Input) *getsale.Output {
	getSale := getsale.GetSale{SaleRepository: s.Repo, UserRepository: s.UserRepository}
	output := getSale.Execute(i)
	return output
}

// Método da controller que comunica com o usecase para a obtenção de dados e pela
// resposta ao cliente
func (s *SaleController) GetSales(i getsales.Input) *getsales.Output {
	getSales := getsales.GetSales{SaleRepository: s.Repo, UserRepository: s.UserRepository}
	output := getSales.Execute(i)
	return output
}

// Método da controller que comunica com o usecase para salvar um objeto de venda
// de acordo com os dados passados no parâmetro
func (s *SaleController) SaveSale(i savesale.Input) *savesale.Output {
	saveSale := savesale.SaveSale{SaleRepository: s.Repo, CommonGlssRepository: s.CommonGlssRepository, PartRepository: s.PartRepository, TemperedGlssRepository: s.TemperedGlssRepository, UserRepository: s.UserRepository}
	output := saveSale.Execute(i)
	return output
}

// Método da controller que comunica com o usacase para atualizar os campos alterados do objeto de
// acordo com os dados passados no parâmetro
func (s *SaleController) PatchSale(i patchsale.Input) *patchsale.Output {
	patchSale := patchsale.PatchSale{SaleRepository: s.Repo, CommonGlssRepository: s.CommonGlssRepository, PartRepository: s.PartRepository, TemperedGlssRepository: s.TemperedGlssRepository, UserRepository: s.UserRepository}
	output := patchSale.Execute(i)
	return output
}

// Método da controller que comunica com o usacase para finalizar a atividade do objeto
// de acordo com os dados passados no parâmetro
func (s *SaleController) CloseSale(i closesale.Input) *closesale.Output {
	closeSale := closesale.CloseSale{SaleRepository: s.Repo, UserRepository: s.UserRepository}
	output := closeSale.Execute(i)
	return output
}

// Método da controller que comunica com o usacase para deletar um objeto
func (s *SaleController) DeleteSale(i deletesale.Input) *deletesale.Output {
	deleteSale := deletesale.DeleteSale{SaleRepository: s.Repo, UserRepository: s.UserRepository}
	output := deleteSale.Execute(i)
	return output
}
