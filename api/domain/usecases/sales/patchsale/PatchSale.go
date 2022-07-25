// Pacote responsável pela o usecase PatchSale que executa
// a ação de receber a confirmação ou cancelamento de produtos
// requeridos por uma venda salvando os status no repositório
package patchsale

import (
	"errors"
	"sync"

	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

type PatchSale struct {
	SaleRepository         repository.SaleRepository
	CommonGlssRepository   repository.CommonGlassRepository
	PartRepository         repository.PartRepository
	TemperedGlssRepository repository.TemperedGlassRepository
}

func (p *PatchSale) Execute(i Input) *Output {

	var (
		wg          sync.WaitGroup
		outPart     *Output
		outTempGlss *Output
		outComnGlss *Output
	)

	if len(i.PartsReq) > 0 {

		wg.Add(1)

		go func() {

			for _, v := range i.PartsReq {

				if v.Id == "" || v.ProductId == "" || v.ProdtQtyReq <= 0 || v.SaleId == "" {
					outPart = new(Output).Init(i.Id, 400, errors.New("some field don't have a value"))
				}

				if v.WasCancelled && v.WasConfirmed {
					outPart = new(Output).Init(i.Id, 400, errors.New("receive a confirmed and cancelled request"))
					break
				}

				if v.WasCancelled {

					err := p.PartRepository.IncreaseQuantity(v.ProductId, v.ProdtQtyReq)
					if err != nil {
						outPart = new(Output).Init(i.Id, 500, err)
						break
					}
				}
			}

			wg.Done()
		}()

	}

	if len(i.CommonGlssReq) > 0 {

		wg.Add(1)

		go func() {
			for _, v := range i.CommonGlssReq {

				if v.Id == "" || v.ProductId == "" || v.ProdtQtyReq == 0 || v.SaleId == "" || v.RequestHeight <= 0 || v.RequestWidth <= 0 {
					outComnGlss = new(Output).Init(i.Id, 400, errors.New("some field don't get a value"))
				}

				if v.WasCancelled && v.WasConfirmed {
					outComnGlss = new(Output).Init(i.Id, 400, errors.New("receive a confirmed and cancelled request"))
					break
				}

				if v.WasCancelled {

					widthRequest := v.RequestWidth * float32(v.ProdtQtyReq)
					heightRequest := v.RequestHeight * float32(v.ProdtQtyReq)

					err := p.CommonGlssRepository.IncreaseArea(v.ProductId, widthRequest, heightRequest)
					if err != nil {
						outComnGlss = new(Output).Init(i.Id, 500, err)
						break
					}

				}

			}

			wg.Done()
		}()

	}

	if len(i.TempGlssReq) > 0 {

		wg.Add(1)

		go func() {
			for _, v := range i.TempGlssReq {

				if v.Id == "" || v.ProductId == "" || v.ProdtQtyReq <= 0 || v.SaleId == "" {
					outPart = new(Output).Init(i.Id, 400, errors.New("some field don't have a value"))
				}

				if v.WasCancelled && v.WasConfirmed {
					outTempGlss = new(Output).Init(i.Id, 400, errors.New("receive a confirmed and cancelled request"))
					break
				}

				if v.WasCancelled {
					err := p.TemperedGlssRepository.IncreaseQuantity(v.ProductId, v.ProdtQtyReq)
					if err != nil {
						outTempGlss = new(Output).Init(i.Id, 500, err)
						break
					}
				}

			}

			wg.Done()
		}()
	}

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

	id, status, err := p.SaleRepository.PatchSale(*i.ConvertToSale())
	return new(Output).Init(id, status, err)
}
