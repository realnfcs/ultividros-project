package repository

import (
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

// Método que pega um vidro temperado no banco de dados de acordo com o id passado
// no parâmetro e o retorna
func (t *TemperedGlassRepositoryMySql) GetTemperedGlass(id string) *entities.TemperedGlass {
	tempGlss := models.ModelTemperedGlass{}

	err := t.GormDb.Where("id = ?", id).First(&tempGlss).Error
	if err != nil {
		return nil
	}

	return tempGlss.TranformToEntity()
}

// Método que pega todos os vidros temperados no banco de dados e o retorna
func (t *TemperedGlassRepositoryMySql) GetTemperedGlasses() *[]entities.TemperedGlass {
	tempGlss := []models.ModelTemperedGlass{}

	err := t.GormDb.Find(&tempGlss).Error
	if err != nil {
		return nil
	}

	return new(models.ModelTemperedGlass).TranformToSliceOfEntity(tempGlss)
}

// Método que salva um vidro temperado no banco de dados de acordo com os dados passados
// no parâmetro
func (t *TemperedGlassRepositoryMySql) SaveTemperedGlass(e entities.TemperedGlass) (string, int, error) {
	tempGlass := new(models.ModelTemperedGlass).TransformToModel(e)

	err := t.GormDb.Create(tempGlass).Error
	if err != nil {
		return "", 401, err
	}

	return tempGlass.ID, 201, nil
}

// Método que atualiza os campos de um vidro temperado de acordo com os dados passados
// no parâmetro
func (t *TemperedGlassRepositoryMySql) UpdateTemperedGlass(e entities.TemperedGlass) (string, int, error) {
	tempGlass := new(models.ModelTemperedGlass).TransformToModel(e)

	err := t.GormDb.Where("id = ?", tempGlass.ID).Omit("created_at").Save(tempGlass).Error
	if err != nil {
		return "", 401, err
	}

	return tempGlass.ID, 401, nil
}

// Método que atualiza os campos de um vidro temperado de acordo com os dados passados
// no parâmetro
func (t *TemperedGlassRepositoryMySql) PatchTemperedGlass(e entities.TemperedGlass) (string, int, error) {
	tempGlass := new(models.ModelTemperedGlass).TransformToModel(e)

	err := t.GormDb.Where("id = ?", tempGlass.ID).Omit("created_at").Updates(tempGlass).Error
	if err != nil {
		return "", 401, err
	}

	return tempGlass.ID, 401, nil
}

// Método que exclui um vidro temperado no banco de dados de acordo com os dados passados
// no parâmetro
func (t *TemperedGlassRepositoryMySql) DeleteTemperedGlass(e entities.TemperedGlass) (int, error) {
	tempGlass := new(models.ModelTemperedGlass).TransformToModel(e)

	err := t.GormDb.Delete(tempGlass).Error
	if err != nil {
		return 401, err
	}
	return 201, nil
}
