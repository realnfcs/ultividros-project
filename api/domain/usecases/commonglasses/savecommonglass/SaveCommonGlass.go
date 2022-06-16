// Pacote responsável pela o usecase SaveTemperedGlass que executa
// a ação de salvamento de um vidro comum e retorna os dado de
// resposta ao cliente
package savecommonglass

import "github.com/realnfcs/ultividros-project/api/domain/repository"

// Usecase responsável pela salvamento do vidro comum em um
// repositório passado por meio da inversão de dados
type SaveCommonGlass struct {
	CommonGlassRepository repository.CommonGlassRepository
}

// Método que executa o procedimento de salvamento do vidro comum
func (s *SaveCommonGlass) Execute(i Input) *Output {
	id, status, err := s.CommonGlassRepository.SaveCommonGlass(*i.ConvertToComnGlss())
	return new(Output).Init(id, status, err)
}
