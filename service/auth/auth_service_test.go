package auth

import (
	"miniproject/entities"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var authService AuthService

type AuthRepoDummy struct{}
type JWTRepoDummy struct{}

func (authRepoDummy AuthRepoDummy) Login(user entities.User) (entities.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost) // Hash password
	return entities.User{
		ID:       1,
		Username: "filipi",
		Email:    "test@gmail.com",
		Password: string(hashedPassword),
		Token:    "321",
	}, nil
}

func (authRepoDummy AuthRepoDummy) FindByUserID(userID int) (entities.User, error) {
	return entities.User{ID: 1,
		Username: "filipi",
		Email:    "test@gmail.com",
		Password: "123",
		Token:    "321",
	}, nil
}

func (authRepoDummy AuthRepoDummy) FindByID(userID int) (entities.User, error) {
	return entities.User{ID: userID, Username: "filipi", Email: "test@gmail.com", Password: "123", Token: "321"}, nil
}

func (authRepoDummy AuthRepoDummy) FindByUserIDs(userID int, users *[]entities.User) error {
	return nil
}

func (authRepoDummy AuthRepoDummy) Register(user entities.User) (entities.User, error) {
	return entities.User{ID: 1,
		Username: "gilang",
		Email:    "test@gmail.com",
		Password: "123",
		Token:    "321",
	}, nil
}

func (jwtRepo JWTRepoDummy) GenerateJWT(userID int) (string, error) {
	return "TokenJWT", nil
}

// Tambahkan implementasi GetUserID pada JWTRepoDummy
func (jwtRepo JWTRepoDummy) GetUserID(next echo.HandlerFunc) echo.HandlerFunc {
	return next
}

func setup() {
	jwtRepo := JWTRepoDummy{}
	authRepoDummy := AuthRepoDummy{}
	authService = *NewAuthService(authRepoDummy, jwtRepo)
}

func TestAuthService_Login(t *testing.T) {
	setup()

	t.Run("sukses login", func(t *testing.T) {
		user, err := authService.Login(entities.User{Email: "test@gmail.com", Password: "123"})
		assert.Nil(t, err)
		assert.Equal(t, 1, user.ID)
		assert.NotEmpty(t, user.Token)
	})

	t.Run("gagal login karena email kosong", func(t *testing.T) {
		user, err := authService.Login(entities.User{Password: "123"})
		assert.NotNil(t, err)
		assert.Equal(t, "email is empty", err.Error())
		assert.Equal(t, 0, user.ID)
	})

	t.Run("gagal login karena password kosong", func(t *testing.T) {
		user, err := authService.Login(entities.User{Email: "test@gmail.com"})
		assert.NotNil(t, err)
		assert.Equal(t, "password is empty", err.Error())
		assert.Equal(t, 0, user.ID)
	})

	t.Run("gagal login karena password salah", func(t *testing.T) {
		user, err := authService.Login(entities.User{Email: "test@gmail.com", Password: "wrongpassword"})
		assert.NotNil(t, err)
		assert.Equal(t, "password is wrong", err.Error())
		assert.Equal(t, 0, user.ID)
	})
}

func TestPasswordHashing(t *testing.T) {
	password := "123"
	hashedPassword, err := HashPassword(password)
	assert.Nil(t, err)
	assert.NotEmpty(t, hashedPassword)

	match := CheckPasswordHash(password, hashedPassword)
	assert.True(t, match)
}

func TestAdd(t *testing.T) {
	t.Run("keduanya positif", func(t *testing.T) {
		result := Add(5, 5)
		if result != 10 {
			t.Error("hasilnya bukan 10")
		}
	})

	t.Run("keduanya negatif", func(t *testing.T) {
		result := Add(-5, -5)
		if result != 0 {
			t.Error("hasilnya bukan 0")
		}
	})
}
