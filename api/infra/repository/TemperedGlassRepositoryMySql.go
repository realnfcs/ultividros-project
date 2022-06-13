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
type TemperedGlassRepositoryMySql struct {
	GormDb *gorm.DB
}

// Struct que auxilia nas querys
type tempdGlssId struct {
	ID string
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

	config, _ := t.GormDb.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	return t, nil
}

// Método para auxiliar nos tests / Método para pegar um id aleatório no repositório (DEMO)
func (t *TemperedGlassRepositoryMySql) GetRandomId() (string, error) {

	tId := tempdGlssId{}

	err := t.GormDb.Model(&models.TemperedGlass{}).Take(&tId).Error
	if err != nil {
		return "", err
	}

	return tId.ID, nil
}

// Método que pega um vidro temperado no banco de dados de acordo com o id passado
// no parâmetro e o retorna
func (t *TemperedGlassRepositoryMySql) GetTemperedGlass(id string) (*entities.TemperedGlass, int, error) {
	tempGlss := new(models.TemperedGlass)

	err := t.GormDb.First(tempGlss, "id = ?", id).Error

	if err != nil {
		return nil, 404, err
	}

	if tempGlss == nil {
		return nil, 500, errors.New("Internal error")
	}

	return tempGlss.TranformToEntity(), 200, nil
}

// Método que pega todos os vidros temperados no banco de dados e o retorna
func (t *TemperedGlassRepositoryMySql) GetTemperedGlasses() (*[]entities.TemperedGlass, int, error) {
	tempGlss := []models.TemperedGlass{}

	err := t.GormDb.Find(&tempGlss).Error
	if err != nil {
		return nil, 500, err
	}

	if len(tempGlss) == 0 {
		return new(models.TemperedGlass).TranformToSliceOfEntity(tempGlss), 404, errors.New("None tempered glasses in server")
	}

	return new(models.TemperedGlass).TranformToSliceOfEntity(tempGlss), 200, nil
}

// Método que salva um vidro temperado no banco de dados de acordo com os dados passados
// no parâmetro
func (t *TemperedGlassRepositoryMySql) SaveTemperedGlass(e entities.TemperedGlass) (string, int, error) {
	tempGlass := new(models.TemperedGlass).TransformToModel(e)

	err := t.GormDb.Create(tempGlass).Error
	if err != nil {
		return "", 400, err
	}

	return tempGlass.ID, 201, nil
}

// Método que atualiza os campos de um vidro temperado de acordo com os dados passados
// no parâmetro
func (t *TemperedGlassRepositoryMySql) UpdateTemperedGlass(e entities.TemperedGlass) (string, int, error) {
	tempGlass := new(models.TemperedGlass).TransformToModel(e)

	id := tempdGlssId{}

	err := t.GormDb.Model(tempGlass).First(&id, "id = ?", tempGlass.ID).Error
	if err != nil {
		return tempGlass.ID, 404, err
	}

	err = t.GormDb.Where("id = ?", tempGlass.ID).Omit("created_at").Save(tempGlass).Error
	if err != nil {
		return "", 400, err
	}

	return tempGlass.ID, 200, nil
}

// Método que atualiza os campos de um vidro temperado de acordo com os dados passados
// no parâmetro
func (t *TemperedGlassRepositoryMySql) PatchTemperedGlass(e entities.TemperedGlass) (string, int, error) {
	tempGlass := new(models.TemperedGlass).TransformToModel(e)

	id := tempdGlssId{}

	err := t.GormDb.Model(tempGlass).First(&id, "id = ?", tempGlass.ID).Error
	if err != nil {
		return tempGlass.ID, 404, err
	}

	err = t.GormDb.Where("id = ?", tempGlass.ID).Omit("created_at").Updates(tempGlass).Error
	if err != nil {
		return "", 400, err
	}

	return tempGlass.ID, 200, nil
}

// Método que exclui um vidro temperado no banco de dados de acordo com os dados passados
// no parâmetro
func (t *TemperedGlassRepositoryMySql) DeleteTemperedGlass(e entities.TemperedGlass) (int, error) {
	tempGlass := new(models.TemperedGlass).TransformToModel(e)

	err := t.GormDb.First(tempGlass, "id = ?", tempGlass.ID).Error
	if err != nil {
		return 404, err
	}

	err = t.GormDb.Delete(tempGlass).Error
	if err != nil {
		return 400, err
	}
	return 200, nil
}
