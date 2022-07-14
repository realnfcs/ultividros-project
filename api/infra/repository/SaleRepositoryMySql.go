package repository

import (
	"errors"
	"time"

	"github.com/realnfcs/ultividros-project/api/domain/entities"
	"github.com/realnfcs/ultividros-project/api/infra/database/models"
	databasemysql "github.com/realnfcs/ultividros-project/api/infra/database/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Struct responsável por armazenar o ponteiro do Gorm que
// faz as querys
type SaleRepositoryMySql struct {
	GormDb *gorm.DB
}

// Struct que auxilia nas querys
type saleId struct {
	ID string
}

var (
	partReqRepo     *PartReqRepositoryMySql
	tempGlssReqRepo *TempGlssReqRepositoryMySql
	comnGlssReqRepo *ComnGlssReqRepository
)

// Método para iniciar o ORM de acordo com a conexão já estabelecida com
// o banco de dados MySQL
func (s *SaleRepositoryMySql) Init() (*SaleRepositoryMySql, error) {
	db, err := new(databasemysql.DatabaseMysql).Init()
	if err != nil {
		return nil, err
	}

	s.GormDb, err = gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	databasemysql.RunMigrations(s.GormDb)

	config, _ := s.GormDb.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	partReqRepo = new(PartReqRepositoryMySql).Init(s.GormDb)
	if partReqRepo == nil {
		return nil, errors.New("Can't inicialize partReq repository")
	}

	tempGlssReqRepo = new(TempGlssReqRepositoryMySql).Init(s.GormDb)
	if tempGlssReqRepo == nil {
		return nil, errors.New("Can't inicialize tempGlssReq repository")
	}

	comnGlssReqRepo = new(ComnGlssReqRepository).Init(s.GormDb)
	if comnGlssReqRepo == nil {
		return nil, errors.New("Can't inicialize comnGlssReq repository")
	}

	return s, nil
}

// CRUD Section //

// Método que pega uma venda e seus produtos requeridos no banco de dados de
// acordo com o id passado no parâmetro e o retorna
func (s *SaleRepositoryMySql) GetSale(id string) (*entities.Sale, int, error) {

	sale := new(models.Sale)
	comnGlssReq := []models.CommonGlssReq{}
	partReq := []models.PartReq{}
	tempGlssReq := []models.TempGlssReq{}

	err := s.GormDb.First(sale, "id = ?", id).Error
	if err != nil {
		return nil, 404, err
	}

	err = s.GormDb.Find(&comnGlssReq, "sale_id = ?", id).Error
	if err != nil {
		return nil, 404, err
	}

	err = s.GormDb.Find(&partReq, "sale_id = ?", id).Error
	if err != nil {
		return nil, 404, err
	}

	err = s.GormDb.Find(&tempGlssReq, "sale_id = ?", id).Error
	if err != nil {
		return nil, 404, err
	}

	sale.CommonGlssReq = comnGlssReq
	sale.PartReq = partReq
	sale.TempGlssReq = tempGlssReq

	return sale.TranformToEntity(), 200, nil
}

// Método que pega todas as vendas no banco de dados e as retorna
func (s *SaleRepositoryMySql) GetSales() (*[]entities.Sale, int, error) {

	sales := []models.Sale{}
	comnGlss := []models.CommonGlssReq{}
	partReq := []models.PartReq{}
	tempGlss := []models.TempGlssReq{}

	err := s.GormDb.Find(&sales).Error
	if err != nil {
		return nil, 500, err
	}

	err = s.GormDb.Find(&comnGlss).Error
	if err != nil {
		return nil, 500, err
	}

	err = s.GormDb.Find(&partReq).Error
	if err != nil {
		return nil, 500, err
	}

	err = s.GormDb.Find(&tempGlss).Error
	if err != nil {
		return nil, 500, err
	}

	for i, v := range sales {
		for _, value := range comnGlss {
			if value.SaleID == v.ID {
				sales[i].CommonGlssReq = append(sales[i].CommonGlssReq, value)
			}
		}

		for _, value := range partReq {
			if value.SaleID == v.ID {
				sales[i].PartReq = append(sales[i].PartReq, value)
			}
		}

		for _, value := range tempGlss {
			if value.SaleID == v.ID {
				sales[i].TempGlssReq = append(sales[i].TempGlssReq, value)
			}
		}
	}

	return new(models.Sale).TransformToSliceOfEntity(sales), 200, nil
}

// Método que salva uma venda no banco de dados de acordo com os dados passados
// no parâmetro
func (s *SaleRepositoryMySql) SaveSale(saleEnt entities.Sale) (string, int, error) {

	sale := new(models.Sale).TransformToModel(saleEnt)

	if sale.ClientID == "" || (len(sale.CommonGlssReq) <= 0 && len(sale.PartReq) <= 0 && len(sale.TempGlssReq) <= 0) {
		return sale.ID, 400, errors.New("Empty field error: some field no got a value")
	}

	err := s.GormDb.Create(sale).Error
	if err != nil {
		return sale.ID, 400, err
	}

	return sale.ID, 201, nil
}
