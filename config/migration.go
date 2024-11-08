package config

import (
	"miniproject/repo/auth"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&auth.User{})
}
