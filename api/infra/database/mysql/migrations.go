package databasemysql

import (
	"github.com/realnfcs/ultividros-project/api/infra/database/models"
	"gorm.io/gorm"
)

// Função que executa as migrations e cria as tabelas no banco de dados
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.TemperedGlass{})
}
