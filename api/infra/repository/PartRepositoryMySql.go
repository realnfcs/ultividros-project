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

// Structs que auxiliam nas querys //
type partId struct {
	ID string
}

type partQty struct {
	Qty uint32
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

// Query Utils Section //

func (p *PartRepositoryMySql) GetPartQuantity(id string) (uint32, error) {

	part := new(models.Part)
	qty := partQty{}

	err := p.GormDb.Model(part).First(&qty, "id = ?", id).Error
	if err != nil {
		return 0, err
	}

	return qty.Qty, nil
}

// CRUD Section //

// Método que pega uma peça no banco de dados de acordo com o id passado
// no parâmetro e o retorna
func (p *PartRepositoryMySql) GetPart(id string) (*entities.Part, int, error) {

	part := new(models.Part)

	err := p.GormDb.First(part, "id = ?", id).Error
	if err != nil {
		return nil, 404, err
	}

	if part == nil {
		return nil, 500, errors.New("Internal error")
	}

	return part.TranformToEntity(), 200, nil
}

// Método que pega todas as peças no banco de dados e as retorna
func (p *PartRepositoryMySql) GetParts() (*[]entities.Part, int, error) {

	parts := []models.Part{}

	err := p.GormDb.Find(&parts).Error
	if err != nil {
		return nil, 500, err
	}

	if len(parts) == 0 {
		return new(models.Part).TranformToSliceOfEntity(parts), 404, errors.New("None parts in server")
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

// Método que atualiza os campos de uma peça de acordo com os dados passados
// no parâmetro
func (p *PartRepositoryMySql) PatchPart(e entities.Part) (string, int, error) {
	part := new(models.Part).TransformToModel(e)

	id := partId{}

	err := p.GormDb.Model(part).First(&id, "id = ?", part.ID).Error
	if err != nil {
		return part.ID, 404, err
	}

	err = p.GormDb.Where("id = ?", part.ID).Omit("created_at").Updates(part).Error
	if err != nil {
		return "", 400, err
	}

	return part.ID, 200, nil
}

// Método que exclui uma peça no banco de dados de acordo com os dados passados
// no parâmetro
func (p *PartRepositoryMySql) DeletePart(e entities.Part) (int, error) {
	part := new(models.Part).TransformToModel(e)

	err := p.GormDb.First(part, "id = ?", part.ID).Error
	if err != nil {
		return 404, err
	}

	err = p.GormDb.Delete(part).Error
	if err != nil {
		return 400, err
	}

	return 200, nil
}
