package repository

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/entities"
	"github.com/realnfcs/ultividros-project/api/infra/database/models"
	"gorm.io/gorm"
)

// Struct responsável por armazenar o ponteiro do Gorm que
// faz as querys
type TempGlssReqRepositoryMySql struct {
	GormDb *gorm.DB
}

// Struct que auxilia nas querys
type tempGlssReqId struct {
	ID string
}

func (p *TempGlssReqRepositoryMySql) Init(g *gorm.DB) *TempGlssReqRepositoryMySql {
	return &TempGlssReqRepositoryMySql{GormDb: g}
}

// Método que pega todos os vidros temperados requeridos pelo cliente em
// uma venda no banco de dados e as retorna
// func (t *TempGlssReqRepositoryMySql) GetTempGlssReqs(string) (*[]entities.TempGlssReq, int, error)

// Método que salva vidros temperados requeridos nas vendas no banco de
// dados de acordo com os dados passados no parâmetro
func (t *TempGlssReqRepositoryMySql) SaveTempGlssReq(e entities.TempGlssReq) (string, int, error) {

	tempGlssReq := new(models.TempGlssReq).TransformToModel(e)

	if tempGlssReq.ProductId == "" || tempGlssReq.ProductName == "" || tempGlssReq.ProductPrice == 0 || tempGlssReq.ProdtQtyReq == 0 || tempGlssReq.SaleID == "" {
		return tempGlssReq.ID, 400, errors.New("Empty field error: some field no got a value")
	}

	err := t.GormDb.Create(tempGlssReq).Error
	if err != nil {
		return tempGlssReq.ID, 400, err
	}

	return tempGlssReq.ID, 201, nil
}
