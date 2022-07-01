// Pacote responsável pela o usecase SaveSale que executa
// a ação de salvamento de uma venda e retorna os dado de
// resposta ao cliente
package savesale

import "github.com/realnfcs/ultividros-project/api/domain/repository"

// Usecase responsável pela salvamento de uma venda em um
// repositório passado por meio da inversão de dados
type SaveSale struct {
	SaleRepository repository.SaleRepository
}

func (s *SaveSale) Execute(i Input) *Output {
	// TODO: Verificar se as quantidades corresponde à quantidade
	// 		 presente no estoque em todos os produtos e se a área
	// 		 requerida nos vidros comuns corresponde e se está disponível
	//		 para a redução.

	id, status, err := s.SaleRepository.SaveSale(*i.ConvertToSale(), *i.ConvertProdInputInProdEnt(), *i.ConvertComnGlssReqInputInComnGlssReqEnt())
	return new(Output).Init(id, status, err)
}
