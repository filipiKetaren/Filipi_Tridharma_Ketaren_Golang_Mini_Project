package auth

import "miniproject/entities"

type AuthInterface interface {
	Login(user entities.User) (entities.User, error)
	Register(user entities.User) (entities.User, error)
	FindUserByIDs(userID int) ([]entities.User, error)
	FindByUserID(userID int) (entities.User, error)
}
