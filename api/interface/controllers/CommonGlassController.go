package controllers

import (
	"github.com/realnfcs/ultividros-project/api/domain/repository"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/commonglasses/getcommonglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/commonglasses/savecommonglass"
)

// Stuct para iniciar a controller dos vidros comuns necessitando
// de um repository para funcionar
type CommonGlassController struct {
	Repo repository.CommonGlassRepository
}

// Método da controller que comunica com o usecase para a obtenção de dados de um
// único vidro comum trazendo a resposta ao cliente
func (c *CommonGlassController) GetCommonGlass(i getcommonglass.Input) *getcommonglass.Output {
	getComnGlss := getcommonglass.GetCommonGlass{CommonGlassRepository: c.Repo}
	output := getComnGlss.Execute(i)
	return output
}

// Método da controller que comunica com o usecase para salvar um objeto de vidro comum
// de acordo com os dados passados no parâmetro
func (c *CommonGlassController) SaveCommonGlass(i savecommonglass.Input) *savecommonglass.Output {
	saveComnGlss := savecommonglass.SaveCommonGlass{CommonGlassRepository: c.Repo}
	output := saveComnGlss.Execute(i)
	return output
}
