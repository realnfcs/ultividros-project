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

func (t *TempGlssReqRepositoryMySql) CancelATempGlssRequest(id string) error {

	tempGlssReq := new(models.TempGlssReq)

	err := t.GormDb.First(tempGlssReq, "id = ?", id).Error
	if err != nil {
		return err
	}

	if tempGlssReq == nil {
		return errors.New("an object in this id don't exist")
	}

	if tempGlssReq.WasConfirmed != false {
		return errors.New("can't cancel a confirmed request of tempered glass")
	}

	err = t.GormDb.Model(tempGlssReq).Where("id = ?", id).Omit("created_at").Update("was_cancelled", true).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *TempGlssReqRepositoryMySql) ConfirmATempGlssRequest(id string) error {

	tempGlssReq := new(models.TempGlssReq)

	err := t.GormDb.First(tempGlssReq, "id = ?", id).Error
	if err != nil {
		return err
	}

	if tempGlssReq == nil {
		return errors.New("an object in this id don't exist")
	}

	if tempGlssReq.WasCancelled != false {
		return errors.New("can't confirm a cancelled request of tempered glass")
	}

	err = t.GormDb.Model(tempGlssReq).Where("id = ?", id).Omit("created_at").Update("was_confirmed", true).Error
	if err != nil {
		return err
	}

	return nil
}

// CRUD Section //

// Método que pega todos os vidros temperados requeridos pelo cliente em
// uma venda no banco de dados e as retorna
func (t *TempGlssReqRepositoryMySql) GetTempGlssReqs(saleId string) (*[]entities.TempGlssReq, error) {

	tempGlssReq := []models.TempGlssReq{}

	err := t.GormDb.Find(&tempGlssReq, "id = ?", saleId).Error
	if err != nil {
		return nil, err
	}

	return new(models.TempGlssReq).TransformToSliceOfEntity(tempGlssReq), nil
}

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
