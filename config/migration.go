package config

import (
	"miniproject/repo/auth"
	"miniproject/repo/plant"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&auth.User{}, &plant.Plant{})
}
