package config

import (
	"miniproject/repo/model"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Plant{}, &model.PlantCondition{})
}
