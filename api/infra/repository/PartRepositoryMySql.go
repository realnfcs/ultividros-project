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
type PartRepositoryMySql struct {
	GormDb *gorm.DB
}

// Struct que auxilia nas querys
type partId struct {
	ID string
}

// Método para iniciar o ORM de acordo com a conexão já estabelecida com
// o banco de dados MySQL
func (p *PartRepositoryMySql) Init() (*PartRepositoryMySql, error) {
	db, err := new(databasemysql.DatabaseMysql).Init()
	if err != nil {
		return nil, err
	}

	p.GormDb, err = gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	databasemysql.RunMigrations(p.GormDb)

	config, _ := p.GormDb.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	return p, nil
}

// Método que pega todas as peças no banco de dados e as retorna
func (p *PartRepositoryMySql) GetParts() (*[]entities.Part, int, error) {

	parts := []models.Part{}

	err := p.GormDb.Find(&parts).Error
	if err != nil {
		return nil, 500, err
	}

	if len(parts) == 0 {
		return new(models.Part).TranformToSliceOfEntity(parts), 404, errors.New("None common glasses in server")
	}

	return new(models.Part).TranformToSliceOfEntity(parts), 200, nil
}

// Método que salva uma peça no banco de dados de acordo com os dados passados
// no parâmetro
func (p *PartRepositoryMySql) SavePart(e entities.Part) (string, int, error) {
	part := new(models.Part).TransformToModel(e)

	if part.Name == "" || part.Price == 0 || part.ForType == "" {
		return part.ID, 400, errors.New("Empty field error: some field no got a value")
	}

	err := p.GormDb.Create(part).Error
	if err != nil {
		return "", 400, err
	}

	return part.ID, 201, nil
}
