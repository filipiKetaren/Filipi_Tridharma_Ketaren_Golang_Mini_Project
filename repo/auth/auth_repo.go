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
	result := authRepo.db.First(&userDB, "email = ?", userDB.Email)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return userDB.ToEntities(), nil
}

func (authRepo AuthRepo) Register(user entities.User) (entities.User, error) {
	userDB := FromEntities(user)
	result := authRepo.db.Create(&userDB)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return userDB.ToEntities(), nil
}

func (authRepo *AuthRepo) FindByUserIDs(userID int, users *[]entities.User) error {
	if err := authRepo.db.Where("id = ?", userID).Find(users).Error; err != nil {
		return err
	}
	return nil
}

func (authRepo *AuthRepo) FindByID(userID int) (entities.User, error) {
	var user entities.User
	err := authRepo.db.First(&user, userID).Error
	if err != nil {
		return entities.User{}, err // Return error jika tidak ditemukan atau query gagal
	}
	return user, nil
}
