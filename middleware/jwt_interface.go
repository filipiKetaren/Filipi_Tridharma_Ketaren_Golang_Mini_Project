package middleware

import "github.com/labstack/echo/v4"

type JwtInterface interface {
	GenerateJWT(userID int) (string, error)
	GetUserID(next echo.HandlerFunc) echo.HandlerFunc
}
