package usecases

import (
	"github.com/NicolasSales0101/ultividros-project/api/domain/entities"
	"github.com/NicolasSales0101/ultividros-project/api/domain/repository"
)

// Usecase responsável pela obtenção um único vidro temperado de um
// repositório voltando um ponteiro da entidade com o mesmo ID passado
// no parâmetro
type GetTemperedGlass struct {
	TemperedGlassRepository repository.TemperedGlassRepository
}

func (g *GetTemperedGlass) Execute(id string) *entities.TemperedGlass {
	return g.TemperedGlassRepository.GetTemperedGlass(id)
}
