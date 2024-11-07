package auth

import "miniproject/entities"

type AuthRepoInterface interface {
	Login(user entities.User) (entities.User, error)
}
