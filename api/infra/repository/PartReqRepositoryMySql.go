package repository

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/entities"
	"github.com/realnfcs/ultividros-project/api/infra/database/models"
	"gorm.io/gorm"
)

// Struct responsável por armazenar o ponteiro do Gorm que
// faz as querys
type PartReqRepositoryMySql struct {
	GormDb *gorm.DB
}

// Struct que auxilia nas querys
type partReqId struct {
	ID string
}

func (p *PartReqRepositoryMySql) Init(g *gorm.DB) *PartReqRepositoryMySql {
	return &PartReqRepositoryMySql{GormDb: g}
}

// Método que pega todas as peças requeridas pelo cliente em uma venda no
// banco de dados e as retorna
// func (p *PartReqRepositoryMySql) GetPartReq(string) (*[]entities.PartsReq, int, error)

// Método que salva peças requeridas nas vendas no banco de dados de acordo
// com os dados passados no parâmetro
func (p *PartReqRepositoryMySql) SavePartReq(e entities.PartsReq) (string, int, error) {

	partReq := new(models.PartReq).TransformToModel(e)

	if partReq.ProductId == "" || partReq.ProductName == "" || partReq.ProductPrice == 0 || partReq.ProdtQtyReq == 0 || partReq.SaleID == "" {
		return partReq.ID, 400, errors.New("Empty field error: some field no got a value")
	}

	err := p.GormDb.Create(partReq).Error
	if err != nil {
		return partReq.ID, 400, err
	}

	return partReq.ID, 201, nil
}
