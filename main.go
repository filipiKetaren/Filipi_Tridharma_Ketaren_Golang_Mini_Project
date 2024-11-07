package main

import (
	"miniproject/config"
	authController "miniproject/controller/auth"
	authRepo "miniproject/repo/auth"
	authService "miniproject/service/auth"

	"github.com/labstack/echo/v4"
)

func main() {
	db, _ := config.ConnectDatabase()
	config.MigrateDB(db)
	e := echo.New()
	authRepo := authRepo.NewAuthRepo(db)
	authService := authService.NewAuthService(authRepo)
	authController := authController.NewAuthController(authService)

	e.POST("/login", authController.LoginController)
	e.Start(":8000")
}
