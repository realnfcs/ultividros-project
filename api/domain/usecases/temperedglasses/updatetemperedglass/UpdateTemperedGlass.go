// Pacote responsável pela o usecase UpdateTemperedGlass que executa
// a ação de atualização de um vidro temperado e retorna os dado de
// resposta ao cliente
package updatetemperedglass

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável pela atualização de um vidro temperado em um
// repositório passado por meio da inversão de dados
type UpdateTemperedGlass struct {
	TemperedGlassRepository repository.TemperedGlassRepository
	UserRepository          repository.UserRepository
}

func (u *UpdateTemperedGlass) Execute(i Input) *Output {

	ocup, err := u.UserRepository.VerifyOccupation(i.UserId)
	if err != nil {
		return new(Output).Init("", 400, err)
	}

	if !ocup {
		return new(Output).Init("", 401, errors.New("unauthorized"))
	}

	id, status, err := u.TemperedGlassRepository.UpdateTemperedGlass(*i.ConvertToTempGlss())
	return new(Output).Init(id, status, err)
}
