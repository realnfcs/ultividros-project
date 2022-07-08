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

// Structs que auxiliam nas querys //
type tempGlssId struct {
	ID string
}

type tempGlssQty struct {
	Quantity uint32
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

// Query Utils Section //

func (t *TemperedGlassRepositoryMySql) GetTempGlssQty(id string) (uint32, error) {

	tempGlss := new(models.TemperedGlass)
	qty := tempGlssQty{}

	err := t.GormDb.Model(tempGlss).First(&qty, "id = ?", id).Error
	if err != nil {
		return 0, err
	}

	return qty.Quantity, nil
}

// Método para auxiliar nos tests / Método para pegar um id aleatório no repositório (DEMO)
func (t *TemperedGlassRepositoryMySql) GetRandomId() (string, error) {

	tId := new(tempGlssId)

	err := t.GormDb.Model(&models.TemperedGlass{}).Take(tId).Error
	if err != nil {
		return "", err
	}

	return tId.ID, nil
}

// Model Utils Section //

// Método que reduz a quantidade em estoque de um vidro temperado identificado pelo
// id passado como parâmetro
func (t *TemperedGlassRepositoryMySql) ReduceQuantity(id string, qtyReq uint32) error {

	if id == "" || qtyReq <= 0 {
		return errors.New("id or qty request don't have a value")
	}

	tempGlss := new(models.TemperedGlass)
	qty := new(tempGlssQty)

	err := t.GormDb.Model(tempGlss).First(qty, "id = ?", id).Error
	if err != nil {
		return err
	}

	newQty := qty.Quantity - qtyReq

	err = t.GormDb.Model(tempGlss).Where("id = ?", id).Omit("created_at").Update("quantity", newQty).Error
	if err != nil {
		return err
	}

	return nil
}

// CRUD Section //

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

	// DEMO
	if tempGlass.Name == "" || tempGlass.Price == 0 || tempGlass.Type == "" || tempGlass.Color == "" || tempGlass.GlassSheets == 0 || tempGlass.Milimeter == 0 || tempGlass.Height == 0 || tempGlass.Width == 0 {
		return tempGlass.ID, 400, errors.New("Empty field error: some field no got a value")
	}

	err := t.GormDb.Create(tempGlass).Error
	if err != nil {
		return "", 400, err
	}

	return tempGlass.ID, 201, nil
}

// Método que atualiza os campos de um vidro temperado de acordo com os dados passados
// no parâmetro
func (t *TemperedGlassRepositoryMySql) UpdateTemperedGlass(e entities.TemperedGlass) (string, int, error) {

	// DEMO
	if e.Name == "" || e.Price == 0 || e.Type == "" || e.Color == "" || e.GlassSheets == 0 || e.Milimeter == 0 || e.Height == 0 || e.Width == 0 {
		return e.Id, 400, errors.New("Empty field error: some field no got a value")
	}

	tempGlass := new(models.TemperedGlass).TransformToModel(e)

	id := tempGlssId{}

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

	id := tempGlssId{}

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
