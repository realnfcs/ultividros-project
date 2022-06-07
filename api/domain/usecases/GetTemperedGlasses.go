package usecases

import (
	"github.com/realnfcs/ultividros-project/api/domain/entities"
	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável pela obtenção de todos os vidros temperados de um
// repositório voltando um ponteiro de array de vidros temperados
type GetTemperedGlasses struct {
	TemperedGlassRepository repository.TemperedGlassRepository
}

// Método que executa o procedimento de pegar os vidros temperados
func (g *GetTemperedGlasses) Execute() *[]entities.TemperedGlass {
	return g.TemperedGlassRepository.GetTemperedGlasses()
}
