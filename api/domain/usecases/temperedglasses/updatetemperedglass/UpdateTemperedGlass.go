package updatetemperedglass

import (
	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável pela atualização de um vidro temperado em um
// repositório passado por meio da inversão de dados
type UpdateTemperedGlass struct {
	TemperedGlassRepository repository.TemperedGlassRepository
}

func (u *UpdateTemperedGlass) Execute(i Input) *Output {
	id, status, err := u.TemperedGlassRepository.UpdateTemperedGlass(*i.ConvertToTempGlss())
	return new(Output).Init(id, status, err)
}
