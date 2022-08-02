package controllers

import (
	"github.com/realnfcs/ultividros-project/api/domain/repository"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/parts/deletepart"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/parts/getpart"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/parts/getparts"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/parts/patchpart"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/parts/savepart"
)

// Stuct para iniciar a controller das peças necessitando
// de um repository para funcionar
type PartController struct {
	Repo           repository.PartRepository
	UserRepository repository.UserRepository
}

// Método da controller que comunica com o usecase para a obtenção de dados de uma
// única peça trazendo a resposta ao cliente
func (p *PartController) GetPart(i getpart.Input) *getpart.Output {
	getPart := getpart.GetPart{PartRepository: p.Repo}
	output := getPart.Execute(i)
	return output
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
	savePart := savepart.SavePart{PartRepository: p.Repo, UserRepository: p.UserRepository}
	output := savePart.Execute(i)
	return output
}

// Método da controller que comunica com o usacase para atualizar os campos alterados do objeto de
// acordo com os dados passados no parâmetro
func (p *PartController) PatchPart(i patchpart.Input) *patchpart.Output {
	patchPart := patchpart.PatchPart{PartRepository: p.Repo, UserRepository: p.UserRepository}
	output := patchPart.Execute(i)
	return output
}

// Método da controller que comunica com o usacase para deletar um objeto
func (p *PartController) DeletePart(i deletepart.Input) *deletepart.Output {
	deletePart := deletepart.DeletePart{PartRepository: p.Repo, UserRepository: p.UserRepository}
	output := deletePart.Execute(i)
	return output
}
