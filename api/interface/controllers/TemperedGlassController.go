package controllers

import (
	"github.com/NicolasSales0101/ultividros-project/api/domain/entities"
	"github.com/NicolasSales0101/ultividros-project/api/domain/repository"
	"github.com/NicolasSales0101/ultividros-project/api/domain/usecases"
)

// Stuct para iniciar a controller dos vidros temperados necessitando
// de um repository para funcionar
type TemperedGlassController struct {
	Repo repository.TemperedGlassRepository
}

// Método da controller que comunica com o usecase para a obtenção de dados de um
// único vidro temperado trazendo a resposta ao cliente
func (t *TemperedGlassController) GetTemperedGlass(id string) *entities.TemperedGlass {
	getTempGlss := usecases.GetTemperedGlass{TemperedGlassRepository: t.Repo}
	temperedGlass := getTempGlss.Execute(id)
	return temperedGlass
}

// Método da controller que comunica com o usecase para a obtenção de dados e pela
// resposta ao cliente
func (t *TemperedGlassController) GetTemperedGlasses() *[]entities.TemperedGlass {
	getTempGlss := usecases.GetTemperedGlasses{TemperedGlassRepository: t.Repo}
	temperedGlasses := getTempGlss.Execute()
	return temperedGlasses
}

// Método da controller que comunica com o usecase para salvar um objeto de acordo
// com os dados passados no parâmetro
func (t *TemperedGlassController) SaveTemperedGlasses(e entities.TemperedGlass) error {
	saveTempGlss := usecases.SaveTemperedGlass{TemperedGlassRepository: t.Repo}

	err := saveTempGlss.Execute(e)
	if err != nil {
		return err
	}

	return nil
}

// Método da controller que comunica com o usacase para atualizar os campos do objeto de
// acordo com os dados passados no parâmetro
func (t *TemperedGlassController) UpdateTemperedGlasses(e entities.TemperedGlass) error {
	updateTempGlss := usecases.UpdateTemperedGlass{TemperedGlassRepository: t.Repo}

	err := updateTempGlss.Execute(e)
	if err != nil {
		return err
	}

	return nil
}
