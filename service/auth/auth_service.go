package auth

import (
	"errors"
	"miniproject/entities"
	"miniproject/repo/auth"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(ar auth.AuthRepoInterface) *AuthService {
	return &AuthService{
		authRepoInterface: ar,
	}
}

type AuthService struct {
	authRepoInterface auth.AuthRepoInterface
}

func (authService AuthService) Login(user entities.User) (entities.User, error) {
	if user.Email == "" {
		return entities.User{}, errors.New("email empty")
	} else if user.Password == "" {
		return entities.User{}, errors.New("password empty")
	}
	var err error
	user, err = authService.authRepoInterface.Login(user)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}
func (authService AuthService) Register(user entities.User) (entities.User, error) {
	return entities.User{}, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
