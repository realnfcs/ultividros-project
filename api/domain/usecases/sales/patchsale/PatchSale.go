// Pacote responsável pela o usecase PatchSale que executa
// a ação de atualizar somente os campos que tiveram mudança de uma
// venda e seus produtos salvando-os no repositório
package patchsale

import (
	"github.com/realnfcs/ultividros-project/api/domain/repository"
)

type PatchSale struct {
	SaleRepository         repository.SaleRepository
	CommonGlssRepository   repository.CommonGlassRepository
	PartRepository         repository.PartRepository
	TemperedGlssRepository repository.TemperedGlassRepository
}

// TODO: Fazer as veirificações se houverão modificações na quantidade requerida
//		 nos produtos e na área dos vidros comuns se tiver no pedido.
func (p *PatchSale) Execute(i Input) *Output
