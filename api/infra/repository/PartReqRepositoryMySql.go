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

// Query Utils Section //

// Struct que auxilia nas querys
type partReqId struct {
	ID string
}

func (p *PartReqRepositoryMySql) Init(g *gorm.DB) *PartReqRepositoryMySql {
	return &PartReqRepositoryMySql{GormDb: g}
}

func (p *PartReqRepositoryMySql) CancelAPartRequest(id string) error {

	partReq := new(models.PartReq)

	err := p.GormDb.First(partReq, "id = ?", id).Error
	if err != nil {
		return err
	}

	if partReq == nil {
		return errors.New("an object in this id don't exist")
	}

	if partReq.WasConfirmed != false {
		return errors.New("can't cancel a confirmed request of part")
	}

	err = p.GormDb.Model(partReq).Where("id = ?", id).Omit("created_at").Update("was_cancelled", true).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *PartReqRepositoryMySql) ConfirmAPartRequest(id string) error {

	partReq := new(models.PartReq)

	err := p.GormDb.First(partReq, "id = ?", id).Error
	if err != nil {
		return err
	}

	if partReq == nil {
		return errors.New("an object in this id don't exist")
	}

	if partReq.WasCancelled != false {
		return errors.New("can't confirm a cancelled request of part")
	}

	err = p.GormDb.Model(partReq).Where("id = ?", id).Omit("created_at").Update("was_confirmed", true).Error
	if err != nil {
		return err
	}

	return nil
}

// CRUD Section //

// Método que pega todas as peças requeridas pelo cliente em uma venda no
// banco de dados e as retorna
func (p *PartReqRepositoryMySql) GetPartReq(saleId string) (*[]entities.PartsReq, error) {

	partReq := []models.PartReq{}

	err := p.GormDb.Find(&partReq, "sale_id = ?", saleId).Error
	if err != nil {
		return nil, err
	}

	return new(models.PartReq).TransformToSliceOfEntity(partReq), nil
}

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
