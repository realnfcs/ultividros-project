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

// Método que pega todas as vendas no banco de dados e as retorna
// func (s *SaleRepositoryMySql) GetSales() (*[]entities.Sale, int, error)

// Método que salva uma venda no banco de dados de acordo com os dados passados
// no parâmetro
func (s *SaleRepositoryMySql) SaveSale(saleEnt entities.Sale /*, partReq []entities.PartsReq, comnGlssReq []entities.CommonGlssReq, tempGlssReq []entities.TempGlssReq*/) (string, int, error) {

	sale := new(models.Sale).TransformToModel(saleEnt)

	if sale.ClientID == "" || (len(sale.CommonGlssReq) <= 0 && len(sale.PartReq) <= 0 && len(sale.TempGlssReq) <= 0) {
		return sale.ID, 400, errors.New("Empty field error: some field no got a value")
	}
	/*
		if len(partReq) > 0 {
			for _, v := range partReq {
				_, status, err := partReqRepo.SavePartReq(v)
				if err != nil {
					return sale.ID, status, err
				}

				// saleEnt.PartsReq[i] = partId
			}
		}

		if len(comnGlssReq) > 0 {
			for _, v := range comnGlssReq {
				_, status, err := comnGlssReqRepo.SaveComnGlssReq(v)
				if err != nil {
					return sale.ID, status, err
				}

				// saleEnt.CommonGlssReqId[i] = comnGlssReqId
			}
		}

		if len(tempGlssReq) > 0 {
			for _, v := range tempGlssReq {
				_, status, err := tempGlssReqRepo.SaveTempGlssReq(v)
				if err != nil {
					return sale.ID, status, err
				}

				// saleEnt.TempGlssReqId[i] = tempGlssReqId
			}
		}
	*/

	err := s.GormDb.Create(sale).Error
	if err != nil {
		return sale.ID, 400, err
	}

	return sale.ID, 201, nil
}
