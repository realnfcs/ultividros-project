package controllers

import (
	"github.com/realnfcs/ultividros-project/api/domain/repository"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/deletetemperedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/gettemperedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/gettemperedglasses"
	patchtemperedglass "github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/patchtempetedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/savetemperedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/updatetemperedglass"
)

// Stuct para iniciar a controller dos vidros temperados necessitando
// de um repository para funcionar
type TemperedGlassController struct {
	Repo           repository.TemperedGlassRepository
	UserRepository repository.UserRepository
}

// Método da controller que comunica com o usecase para a obtenção de dados de um
// único vidro temperado trazendo a resposta ao cliente
func (t *TemperedGlassController) GetTemperedGlass(i gettemperedglass.Input) *gettemperedglass.Output {
	getTempGlss := gettemperedglass.GetTemperedGlass{TemperedGlassRepository: t.Repo}
	output := getTempGlss.Execute(i)
	return output
}

// Método da controller que comunica com o usecase para a obtenção de dados e pela
// resposta ao cliente
func (t *TemperedGlassController) GetTemperedGlasses() *gettemperedglasses.Output {
	getTempGlss := gettemperedglasses.GetTemperedGlasses{TemperedGlassRepository: t.Repo}
	output := getTempGlss.Execute()
	return output
}

// Método da controller que comunica com o usecase para salvar um objeto de vidro temperado
// de acordo com os dados passados no parâmetro
func (t *TemperedGlassController) SaveTemperedGlasses(i savetemperedglass.Input) *savetemperedglass.Output {
	saveTempGlss := savetemperedglass.SaveTemperedGlass{TemperedGlassRepository: t.Repo, UserRepository: t.UserRepository}
	output := saveTempGlss.Execute(i)
	return output
}

// Método da controller que comunica com o usacase para atualizar os campos do objeto de
// acordo com os dados passados no parâmetro
func (t *TemperedGlassController) UpdateTemperedGlasses(i updatetemperedglass.Input) *updatetemperedglass.Output {
	updateTempGlss := updatetemperedglass.UpdateTemperedGlass{TemperedGlassRepository: t.Repo, UserRepository: t.UserRepository}
	output := updateTempGlss.Execute(i)
	return output
}

// Método da controller que comunica com o usacase para atualizar os campos alterados do objeto de
// acordo com os dados passados no parâmetro
func (t *TemperedGlassController) PatchTemperedGlasses(i patchtemperedglass.Input) *patchtemperedglass.Output {
	patchTempGlss := patchtemperedglass.PatchTemperedGlass{TemperedGlassRepository: t.Repo, UserRepository: t.UserRepository}
	output := patchTempGlss.Execute(i)
	return output
}

// Método da controller que comunica com o usacase para deletar um objeto
func (t *TemperedGlassController) DeleteTemperedGlass(i deletetemperedglass.Input) *deletetemperedglass.Output {
	deleteTempGlss := deletetemperedglass.DeleteTemperedGlass{TemperedGlassRepository: t.Repo, UserRepository: t.UserRepository}
	output := deleteTempGlss.Execute(i)
	return output
}
