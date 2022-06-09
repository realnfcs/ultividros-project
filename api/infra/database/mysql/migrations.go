package databasemysql

import (
	"github.com/realnfcs/ultividros-project/api/interface/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(model.ModelTemperedGlass{})
}
