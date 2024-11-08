package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtAlta struct{}

type jwtCustomClaims struct {
	Name   string `json:"name"`
	UserID int    `json:"userID"`
	jwt.RegisteredClaims
}

func (jwtAlta JwtAlta) GenerateJWT(userID int, name string) (string, error) {
	claims := &jwtCustomClaims{
		name,
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}
