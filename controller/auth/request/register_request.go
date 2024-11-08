package request

import "miniproject/entities"

type RegisterRequest struct {
	ID       int    `json: "id"`
	Username string `json: "username"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

func (registerRequest RegisterRequest) ToEntities() entities.User {
	return entities.User{
		ID:       registerRequest.ID,
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}
}
