package response

import "miniproject/entities"

type RegisterResponse struct {
	ID       int    `json: "id"`
	Username string `json: "username"`
	Email    string `json: "email"`
}

func RegisterFromEntities(user entities.User) RegisterResponse {
	return RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}
