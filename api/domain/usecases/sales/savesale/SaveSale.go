// Pacote responsável pela o usecase SaveSale que executa
// a ação de salvamento de uma venda e retorna os dado de
// resposta ao cliente
package savesale

import (
	"errors"
	"fmt"
	"sync"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

// Usecase responsável pela salvamento de uma venda em um
// repositório passado por meio da inversão de dados
type SaveSale struct {
	SaleRepository         repository.SaleRepository
	CommonGlssRepository   repository.CommonGlassRepository
	PartRepository         repository.PartRepository
	TemperedGlssRepository repository.TemperedGlassRepository
}

func (s *SaveSale) Execute(i Input) *Output {
	// Verificação se as quantidades corresponde à quantidade
	// presente no estoque em todos os produtos e se a área
	// requerida nos vidros comuns corresponde e se está disponível
	// para a redução.

	var (
		wg          sync.WaitGroup
		outPart     *Output
		outTempGlss *Output
		outComnGlss *Output
	)

	wg.Add(1)

	go func(output *Output) *Output {
		if len(i.Parts) != 0 {
			for _, v := range i.Parts {
				qty, err := s.PartRepository.GetPartQuantity(v.Id)
				if err != nil {
					output := new(Output).Init("", 400, err)
					return output
				}

				if v.ProdtQtyReq > qty {
					output := new(Output).Init("", 400, errors.New("Part quantity request as big than quantity in stock"))
					return output
				}
			}
		}

		return nil
	}(outPart)

	wg.Add(1)

	go func(output *Output) *Output {
		if len(i.TempGlss) != 0 {
			for _, v := range i.TempGlss {
				qty, err := s.TemperedGlssRepository.GetTempGlssQty(v.Id)
				if err != nil {
					return new(Output).Init("", 400, err)
				}

				if v.ProdtQtyReq > qty {
					return new(Output).Init("", 400, errors.New("Tempered glass quantity request as big than quantity in stock"))
				}
			}
		}

		return nil
	}(outTempGlss)

	wg.Add(1)

	go func(output *Output) *Output {
		if len(i.CommonGlss) != 0 {
			for _, v := range i.CommonGlss {
				area, err := s.CommonGlssRepository.GetArea(v.Id)
				if err != nil {
					return new(Output).Init("", 400, err)
				}

				if v.RequestWidth > area["width"] {
					return new(Output).Init("", 400, errors.New("Common glass width request as big than width in stock"))
				}

				if v.RequestHeight > area["height"] {
					return new(Output).Init("", 400, errors.New("Common glass height request as big than height in stock"))
				}

				glassSheetsQtyReq := v.RequestWidth * v.RequestHeight
				glassSheetsTotalArea := area["width"] * area["height"]

				if glassSheetsQtyReq*float32(v.ProdtQtyReq) > glassSheetsTotalArea {
					return new(Output).Init("", 400, errors.New("Total of quantity area in request in big than area in stock"))
				}
			}
		}

		return nil
	}(outComnGlss)

	wg.Wait()

	if outPart != nil || outTempGlss != nil || outComnGlss != nil {
		allErros := fmt.Sprintf("1. %v\n2. %v\n3. %v", outPart.Error, outTempGlss.Error, outComnGlss.Error)
		return new(Output).Init("", 400, errors.New(allErros))
	}

	id, status, err := s.SaleRepository.SaveSale(*i.ConvertToSale(), *i.ConvertPartReqInputInEnt(), *i.ConvertComnGlssReqInputInEnt(), *i.ConvertTempGlssReqInputInEnt())
	return new(Output).Init(id, status, err)
}
