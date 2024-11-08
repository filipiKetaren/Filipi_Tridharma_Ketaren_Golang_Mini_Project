package request

import "miniproject/entities"

type LoginRequest struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

func (loginReguest LoginRequest) ToEntities() entities.User {
	return entities.User{
		Email:    loginReguest.Email,
		Password: loginReguest.Password,
	}
}
