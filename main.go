package main

import (
	"log"
	"miniproject/config"
	authController "miniproject/controller/auth"
	authRepo "miniproject/repo/auth"
	authService "miniproject/service/auth"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	db, _ := config.ConnectDatabase()
	config.MigrateDB(db)
	e := echo.New()
	authRepo := authRepo.NewAuthRepo(db)
	authService := authService.NewAuthService(authRepo)
	authController := authController.NewAuthController(authService)

	e.POST("/login", authController.LoginController)
	e.Start(":8000")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("failed load env")
	}
}
