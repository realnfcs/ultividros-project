package repository

import (
	"time"

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
