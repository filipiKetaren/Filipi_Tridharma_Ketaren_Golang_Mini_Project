package auth

import (
	"miniproject/entities"
)

type user struct {
	ID       int `gorm: primarykey`
	Username string
	Email    string
	Password string
}

func FromEntities(users entities.User) user {
	return user{
		ID:       users.ID,
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}
}

func (user user) ToEntities() entities.User {
	return entities.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (user user) ToEntitiesNoPassword() entities.User {
	return entities.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}
