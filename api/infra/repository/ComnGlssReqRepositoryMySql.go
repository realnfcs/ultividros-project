package repository

import (
	"errors"

	"github.com/realnfcs/ultividros-project/api/domain/entities"
	"github.com/realnfcs/ultividros-project/api/infra/database/models"
	"gorm.io/gorm"
)

// Struct responsável por armazenar o ponteiro do Gorm que
// faz as querys
type ComnGlssReqRepository struct {
	GormDb *gorm.DB
}

// Struct que auxilia nas querys
type comnGlssReqId struct {
	ID string
}

func (p *ComnGlssReqRepository) Init(g *gorm.DB) *ComnGlssReqRepository {
	return &ComnGlssReqRepository{GormDb: g}
}

// Método que pega todos os vidros comuns requeridos pelo cliente em
// uma venda no banco de dados e as retorna
// func (c *ComnGlssReqRepository) GetComnGlssReqs(string) (*[]entities.TempGlssReq, int, error)

// Método que salva vidros comuns requeridos nas vendas no banco de
// dados de acordo com os dados passados no parâmetro
func (c *ComnGlssReqRepository) SaveComnGlssReq(e entities.CommonGlssReq) (string, int, error) {

	comnGlssReq := new(models.CommonGlssReq).TransformToModel(e)

	if comnGlssReq.ProductId == "" || comnGlssReq.ProductName == "" || comnGlssReq.ProductPrice == 0 || comnGlssReq.ProdtQtyReq == 0 || comnGlssReq.RequestWidth == 0 || comnGlssReq.RequestHeight == 0 || comnGlssReq.SaleID == "" {
		return comnGlssReq.ID, 400, errors.New("Empty field error: some field no got a value")
	}

	err := c.GormDb.Create(comnGlssReq).Error
	if err != nil {
		return comnGlssReq.ID, 400, err
	}

	return comnGlssReq.ID, 201, nil
}
