package response

import "miniproject/entities"

type AuthResponse struct {
	ID       int    `json: "id"`
	Username string `json: "username"`
	Email    string `json: "email"`
	Token    string `json: "token"`
}

func FromEntities(user entities.User) AuthResponse {
	return AuthResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    user.Token,
	}
}
