// Pacote responsável pelos controllers que, neste caso, se comunicam
// com os Input e Output ports dos usecases e se encontra na camada
// Interface and Adapters Layer
package controllers

import "github.com/realnfcs/ultividros-project/api/domain/repository"

// Struct responsável pelo armazenamento organizado de todas as controllers
type Controllers struct {
	TemperedController TemperedGlassController
	CommonController   CommonGlassController
	PartController     PartController
	UserController     UserController
	SaleController     SaleController
}

func (c *Controllers) Init(rt repository.TemperedGlassRepository, rc repository.CommonGlassRepository, rp repository.PartRepository, ru repository.UserRepository, rs repository.SaleRepository) *Controllers {
	return &Controllers{
		TemperedController: TemperedGlassController{Repo: rt, UserRepository: ru},
		CommonController:   CommonGlassController{Repo: rc, UserRepository: ru},
		PartController:     PartController{Repo: rp, UserRepository: ru},
		UserController:     UserController{Repo: ru},
		SaleController:     SaleController{Repo: rs, CommonGlssRepository: rc, PartRepository: rp, TemperedGlssRepository: rt},
	}
}
