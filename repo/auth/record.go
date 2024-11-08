package auth

import "miniproject/entities"

type User struct {
	ID       int `gorm: primarykey`
	Username string
	Email    string
	Password string
}

func FromEntities(user entities.User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (user User) ToEntities() entities.User {
	return entities.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}
