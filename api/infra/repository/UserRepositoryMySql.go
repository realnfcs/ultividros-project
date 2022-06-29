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
type UserRepositoryMySql struct {
	GormDb *gorm.DB
}

// Struct que auxilia nas querys
type userId struct {
	ID string
}

// Método para iniciar o ORM de acordo com a conexão já estabelecida com
// o banco de dados MySQL
func (u *UserRepositoryMySql) Init() (*UserRepositoryMySql, error) {
	db, err := new(databasemysql.DatabaseMysql).Init()
	if err != nil {
		return nil, err
	}

	u.GormDb, err = gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	databasemysql.RunMigrations(u.GormDb)

	config, _ := u.GormDb.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	return u, nil
}

// Método que pega um usuário no banco de dados de acordo com o id passado
// no parâmetro e o retorna
func (u *UserRepositoryMySql) GetUser(id string) (*entities.User, int, error) {

	user := new(models.User)

	err := u.GormDb.First(user, "id = ?", id).Error
	if err != nil {
		return nil, 404, err
	}

	if user == nil {
		return nil, 500, errors.New("Internal error")
	}

	return user.TranformToEntity(), 200, nil
}

// Método que pega todos os usuários no banco de dados e as retorna
func (u *UserRepositoryMySql) GetUsers() (*[]entities.User, int, error) {

	users := []models.User{}

	err := u.GormDb.Find(&users).Error
	if err != nil {
		return nil, 500, err
	}

	if len(users) == 0 {
		return new(models.User).TranformToSliceOfEntity(users), 404, errors.New("None users in server")
	}

	return new(models.User).TranformToSliceOfEntity(users), 200, nil
}

// Método que salva um usuário no banco de dados de acordo com os dados passados
// no parâmetro
func (u *UserRepositoryMySql) SaveUser(e entities.User) (string, int, error) {

	user := new(models.User).TransformToModel(e)

	if user.Name == "" || user.Email == "" || user.Password == "" || user.Occupation == "" {
		return user.ID, 400, errors.New("Empty field error: some field no got a value")
	}

	err := u.GormDb.Create(user).Error
	if err != nil {
		return user.ID, 400, err
	}

	return user.ID, 201, nil
}

// Método que atualiza os campos de um usuário de acordo com os dados passados
// no parâmetro
func (u *UserRepositoryMySql) PatchUser(e entities.User) (string, int, error) {

	user := new(models.User).TransformToModel(e)

	id := userId{}

	err := u.GormDb.Model(user).First(&id, "id = ?", user.ID).Error
	if err != nil {
		return user.ID, 404, err
	}

	err = u.GormDb.Where("id = ?", user.ID).Omit("created_at").Updates(user).Error
	if err != nil {
		return user.ID, 400, err
	}

	return user.ID, 200, nil
}

// Método que exclui um usuário no banco de dados de acordo com os dados passados
// no parâmetro
func (u *UserRepositoryMySql) DeleteUser(e entities.User) (int, error) {

	user := new(models.User).TransformToModel(e)

	err := u.GormDb.First(user, "id = ?", user.ID).Error
	if err != nil {
		return 404, err
	}

	err = u.GormDb.Delete(user).Error
	if err != nil {
		return 400, err
	}

	return 200, nil
}
