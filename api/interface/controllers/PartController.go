package controllers

import (
	"github.com/realnfcs/ultividros-project/api/domain/repository"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/parts/getparts"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/parts/savepart"
)

// Stuct para iniciar a controller das peças necessitando
// de um repository para funcionar
type PartController struct {
	Repo repository.PartRepository
}

// Método da controller que comunica com o usecase para a obtenção de dados e pela
// resposta ao cliente
func (p *PartController) GetParts() *getparts.Output {
	getParts := getparts.GetParts{PartRepository: p.Repo}
	output := getParts.Execute()
	return output
}

// Método da controller que comunica com o usecase para salvar um objeto peça
// de acordo com os dados passados no parâmetro
func (p *PartController) SavePart(i savepart.Input) *savepart.Output {
	savePart := savepart.SavePart{PartRepository: p.Repo}
	output := savePart.Execute(i)
	return output
}
