package controllers

import (
	"github.com/realnfcs/ultividros-project/api/domain/repository"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/commonglasses/deletecommonglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/commonglasses/getcommonglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/commonglasses/getcommonglasses"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/commonglasses/patchcommonglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/commonglasses/savecommonglass"
)

// Stuct para iniciar a controller dos vidros comuns necessitando
// de um repository para funcionar
type CommonGlassController struct {
	Repo           repository.CommonGlassRepository
	UserRepository repository.UserRepository
}

// Método da controller que comunica com o usecase para a obtenção de dados de um
// único vidro comum trazendo a resposta ao cliente
func (c *CommonGlassController) GetCommonGlass(i getcommonglass.Input) *getcommonglass.Output {
	getComnGlss := getcommonglass.GetCommonGlass{CommonGlassRepository: c.Repo}
	output := getComnGlss.Execute(i)
	return output
}

// Método da controller que comunica com o usecase para a obtenção de dados e pela
// resposta ao cliente
func (c *CommonGlassController) GetCommonGlasses() *getcommonglasses.Output {
	getComnGlss := getcommonglasses.GetCommonGlasses{CommonGlassRepository: c.Repo}
	output := getComnGlss.Execute()
	return output
}

// Método da controller que comunica com o usecase para salvar um objeto de vidro comum
// de acordo com os dados passados no parâmetro
func (c *CommonGlassController) SaveCommonGlass(i savecommonglass.Input) *savecommonglass.Output {
	saveComnGlss := savecommonglass.SaveCommonGlass{CommonGlassRepository: c.Repo, UserRepository: c.UserRepository}
	output := saveComnGlss.Execute(i)
	return output
}

// Método da controller que comunica com o usacase para atualizar os campos alterados do objeto de
// acordo com os dados passados no parâmetro
func (c *CommonGlassController) PatchCommonGlass(i patchcommonglass.Input) *patchcommonglass.Output {
	patchComnGlss := patchcommonglass.PatchCommonGlass{CommonGlassRepository: c.Repo, UserRepository: c.UserRepository}
	output := patchComnGlss.Execute(i)
	return output
}

// Método da controller que comunica com o usacase para deletar um objeto
func (c *CommonGlassController) DeleteCommonGlass(i deletecommonglass.Input) *deletecommonglass.Output {
	deleteComnGlss := deletecommonglass.DeleteCommonGlass{CommonGlassRepository: c.Repo, UserRepository: c.UserRepository}
	output := deleteComnGlss.Execute(i)
	return output
}
