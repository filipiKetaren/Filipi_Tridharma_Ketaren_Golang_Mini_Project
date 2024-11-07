package auth

import (
	"miniproject/entities"

	"gorm.io/gorm"
)

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

type AuthRepo struct {
	db *gorm.DB
}

func (authRepo AuthRepo) Login(user entities.User) (entities.User, error) {
	userDB := FromEntities(user)
	result := authRepo.db.First(&userDB, "email = ? AND password = ?", userDB.Email, userDB.Password)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return userDB.ToEntities(), nil
}
