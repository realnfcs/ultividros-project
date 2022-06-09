package controllers

import (
	"github.com/realnfcs/ultividros-project/api/domain/repository"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/gettemperedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/gettemperedglasses"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/savetemperedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/updatetemperedglass"
)

// TODO: Insert Input and Output ports in each controller's functions

// Stuct para iniciar a controller dos vidros temperados necessitando
// de um repository para funcionar
type TemperedGlassController struct {
	Repo repository.TemperedGlassRepository
}

// Método da controller que comunica com o usecase para a obtenção de dados de um
// único vidro temperado trazendo a resposta ao cliente
func (t *TemperedGlassController) GetTemperedGlass(i gettemperedglass.Input) *gettemperedglass.Output {
	getTempGlss := gettemperedglass.GetTemperedGlass{TemperedGlassRepository: t.Repo}
	temperedGlass := getTempGlss.Execute(i)
	return temperedGlass
}

// Método da controller que comunica com o usecase para a obtenção de dados e pela
// resposta ao cliente
func (t *TemperedGlassController) GetTemperedGlasses() *gettemperedglasses.Output {
	getTempGlss := gettemperedglasses.GetTemperedGlasses{TemperedGlassRepository: t.Repo}
	temperedGlasses := getTempGlss.Execute()
	return temperedGlasses
}

// Método da controller que comunica com o usecase para salvar um objeto de acordo
// com os dados passados no parâmetro
func (t *TemperedGlassController) SaveTemperedGlasses(i savetemperedglass.Input) *savetemperedglass.Output {
	saveTempGlss := savetemperedglass.SaveTemperedGlass{TemperedGlassRepository: t.Repo}
	output := saveTempGlss.Execute(i)
	return output
}

// Método da controller que comunica com o usacase para atualizar os campos do objeto de
// acordo com os dados passados no parâmetro
func (t *TemperedGlassController) UpdateTemperedGlasses(i updatetemperedglass.Input) *updatetemperedglass.Output {
	updateTempGlss := updatetemperedglass.UpdateTemperedGlass{TemperedGlassRepository: t.Repo}

	err := updateTempGlss.Execute(i)
	if err != nil {
		return err
	}

	return nil
}
