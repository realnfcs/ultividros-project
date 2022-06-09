package repository

import (
	"github.com/realnfcs/ultividros-project/api/domain/entities"
	databasemysql "github.com/realnfcs/ultividros-project/api/infra/database/mysql"
	"github.com/realnfcs/ultividros-project/api/interface/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Struct responsável por armazenar o ponteiro do Gorm que
// faz as querys
type TemperedGlassRepositoryMySql struct {
	GormDb *gorm.DB
}

// Método para iniciar o ORM de acordo com a conexão já estabelecida com
// o banco de dados MySQL
func (t *TemperedGlassRepositoryMySql) Init() (*TemperedGlassRepositoryMySql, error) {
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

	return t, nil
}

func (t *TemperedGlassRepositoryMySql) GetTemperedGlass(id string) *entities.TemperedGlass {
	tempGlss := model.ModelTemperedGlass{}

	err := t.GormDb.First(&tempGlss, id).Error
	if err != nil {
		return nil
	}

	return tempGlss.TranformToEntity()
}

func (t *TemperedGlassRepositoryMySql) GetTemperedGlasses() *[]entities.TemperedGlass {
	tempGlss := []model.ModelTemperedGlass{}

	err := t.GormDb.Find(&tempGlss).Error
	if err != nil {
		return nil
	}

	return new(model.ModelTemperedGlass).TranformToSliceOfEntity(tempGlss)
}

func (t *TemperedGlassRepositoryMySql) SaveTemperedGlass(e entities.TemperedGlass) (string, int, error) {
	tempGlass := new(model.ModelTemperedGlass).TransformToModel(e)

	err := t.GormDb.Create(tempGlass).Error
	if err != nil {
		return "", 401, err
	}

	return tempGlass.ID, 201, nil
}

func (t *TemperedGlassRepositoryMySql) UpdateTemperedGlass(e entities.TemperedGlass) (string, int, error) {
	tempGlass := new(model.ModelTemperedGlass).TransformToModel(e)

	err := t.GormDb.Save(tempGlass).Error
	if err != nil {
		return "", 401, err
	}

	return tempGlass.ID, 401, nil
}
