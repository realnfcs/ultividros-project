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
type CommonGlassRepositoryMySql struct {
	GormDb *gorm.DB
}

// Método para iniciar o ORM de acordo com a conexão já estabelecida com
// o banco de dados MySQL
func (t *CommonGlassRepositoryMySql) Init() (*CommonGlassRepositoryMySql, error) {
	db, err := new(databasemysql.DatabaseMysql).Init()
	if err != nil {
		return nil, err
	}

	t.GormDb, err = gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	databasemysql.RunMigrations(t.GormDb)

	config, _ := t.GormDb.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	return t, nil
}

// Método que pega um vidro comum no banco de dados de acordo com o id passado
// no parâmetro e o retorna
func (c *CommonGlassRepositoryMySql) GetCommonGlass(id string) (*entities.CommonGlass, int, error) {
	commonGlss := new(models.CommonGlass)

	err := c.GormDb.First(commonGlss, "id = ?", id).Error
	if err != nil {
		return nil, 404, err
	}

	if commonGlss == nil {
		return nil, 500, errors.New("Internal error")
	}

	return commonGlss.TranformToEntity(), 200, nil
}

// Método que salva um vidro comum no banco de dados de acordo com os dados passados
// no parâmetro
func (c *CommonGlassRepositoryMySql) SaveCommonGlass(e entities.CommonGlass) (string, int, error) {
	commonGlss := new(models.CommonGlass).TransformToModel(e)

	// DEMO
	if commonGlss.Name == "" || commonGlss.Price == 0 || commonGlss.Type == "" || commonGlss.Color == "" || commonGlss.Milimeter == 0 || commonGlss.HeightAvailable == 0 || commonGlss.WidthAvailable == 0 {
		return commonGlss.ID, 400, errors.New("Empty field error: some field no got a value")
	}

	err := c.GormDb.Create(commonGlss).Error
	if err != nil {
		return "", 400, err
	}

	return commonGlss.ID, 201, nil
}
