package auth

import "miniproject/entities"

type AuthRepoInterface interface {
	Login(user entities.User) (entities.User, error)
	Register(user entities.User) (entities.User, error)
	FindByUserIDs(userID int, users *[]entities.User) error
	FindByID(userID int) (entities.User, error)
}
