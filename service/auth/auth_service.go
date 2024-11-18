package auth

import (
	"miniproject/constant"
	"miniproject/entities"
	"miniproject/middleware"
	"miniproject/repo/auth"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(ar auth.AuthRepoInterface, jt middleware.JwtInterface) *AuthService {
	return &AuthService{
		authRepoInterface: ar,
		jwtInterface:      jt,
	}
}

type AuthService struct {
	authRepoInterface auth.AuthRepoInterface
	jwtInterface      middleware.JwtInterface
}

func (authService AuthService) Login(user entities.User) (entities.User, error) {
	if user.Email == "" {
		return entities.User{}, constant.EMAIL_IS_EMPTY
	} else if user.Password == "" {
		return entities.User{}, constant.PASSWORD_IS_EMPTY
	}

	oldPassword := user.Password

	user, err := authService.authRepoInterface.Login(user)
	if err != nil {
		return entities.User{}, err
	}

	match := CheckPasswordHash(oldPassword, user.Password)
	if !match {
		return entities.User{}, constant.PASSWORD_IS_WRONG
	}

	token, err := authService.jwtInterface.GenerateJWT(user.ID)
	if err != nil {
		return entities.User{}, err
	}

	user.Token = token
	return user, nil
}
func (authService AuthService) Register(user entities.User) (entities.User, error) {
	if user.Email == "" {
		return entities.User{}, constant.EMAIL_IS_EMPTY
	} else if user.Password == "" {
		return entities.User{}, constant.PASSWORD_IS_EMPTY
	}

	hash, _ := HashPassword(user.Password)
	user.Password = hash
	user, err := authService.authRepoInterface.Register(user)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (authService AuthService) FindUserByIDs(userID int) ([]entities.User, error) {
	var users []entities.User
	err := authService.authRepoInterface.FindByUserIDs(userID, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// FindByUserID mencari user berdasarkan ID
func (authService AuthService) FindByUserID(userID int) (entities.User, error) {
	user, err := authService.authRepoInterface.FindByID(userID)
	if err != nil {
		return entities.User{}, err // Return empty user struct jika error
	}
	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
