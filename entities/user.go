package entities

type User struct {
	ID       int
	Username string
	Email    string
	Password string
	Token    string
}

type UserResponse struct {
	ID       int
	Username string
	Email    string
}
