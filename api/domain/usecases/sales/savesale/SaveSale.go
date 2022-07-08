// Pacote responsável pela o usecase SaveSale que executa
// a ação de salvamento de uma venda e retorna os dado de
// resposta ao cliente
package savesale

import (
	"errors"
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

	go func() {

		if len(i.Parts) > 0 {

			for _, v := range i.Parts {
				qty, err := s.PartRepository.GetPartQuantity(v.ProductId)

				if err != nil {
					outPart = new(Output).Init("", 400, err)
					wg.Done()
					return
				}

				if v.ProdQtyReq > qty {
					outPart = new(Output).Init("", 400, errors.New("Part quantity request as big than quantity in stock"))
					wg.Done()
					return
				}

				err = s.PartRepository.ReduceQuantity(v.ProductId, v.ProdQtyReq)
				if err != nil {
					outPart = new(Output).Init("", 500, err)
					wg.Done()
					return
				}
			}
		}

		wg.Done()

		return

	}()

	wg.Add(1)

	go func() {

		if len(i.TempGlss) > 0 {
			for _, v := range i.TempGlss {
				qty, err := s.TemperedGlssRepository.GetTempGlssQty(v.ProductId)
				if err != nil {
					outTempGlss = new(Output).Init("", 400, err)
					wg.Done()
					return
				}

				if v.ProdQtyReq > qty {
					outTempGlss = new(Output).Init("", 400, errors.New("Tempered glass quantity request as big than quantity in stock"))
					wg.Done()
					return
				}

				err = s.TemperedGlssRepository.ReduceQuantity(v.ProductId, v.ProdQtyReq)
				if err != nil {
					outTempGlss = new(Output).Init("", 500, err)
					wg.Done()
					return
				}
			}

		}

		wg.Done()

		return

	}()

	wg.Add(1)

	go func() {
		if len(i.CommonGlss) > 0 {

			for _, v := range i.CommonGlss {

				area, err := s.CommonGlssRepository.GetArea(v.ProductId)

				if err != nil {
					outComnGlss = new(Output).Init("", 400, err)
					wg.Done()
					return
				}

				if v.RequestWidth > area["width"] || v.RequestHeight > area["height"] {
					outComnGlss = new(Output).Init("", 400, errors.New("Common glass width or height request as big than area in stock"))
					wg.Done()
					return
				}

				glassSheetsQtyReq := v.RequestWidth * v.RequestHeight
				glassSheetsTotalArea := area["width"] * area["height"]

				if glassSheetsQtyReq*float32(v.ProdQtyReq) > glassSheetsTotalArea {
					outComnGlss = new(Output).Init("", 400, errors.New("Total of quantity area in request in big than area in stock"))
					wg.Done()
					return
				}

				err = s.CommonGlssRepository.ReduceArea(v.ProductId, v.RequestWidth, v.RequestHeight)
				if err != nil {
					outComnGlss = new(Output).Init("", 500, err)
					wg.Done()
					return
				}
			}
		}

		wg.Done()
		return
	}()

	wg.Wait()

	if outPart != nil {
		return outPart
	}

	if outComnGlss != nil {
		return outComnGlss
	}

	if outTempGlss != nil {
		return outTempGlss
	}

	id, status, err := s.SaleRepository.SaveSale(*i.ConvertToSale())
	return new(Output).Init(id, status, err)
}
