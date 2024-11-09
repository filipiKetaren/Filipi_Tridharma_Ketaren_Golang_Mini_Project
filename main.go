package main

import (
	"log"
	"miniproject/config"
	authController "miniproject/controller/auth"
	plantController "miniproject/controller/plant"
	"miniproject/middleware"
	authRepo "miniproject/repo/auth"
	plantRepo "miniproject/repo/plant"
	"miniproject/route"
	authService "miniproject/service/auth"
	plantService "miniproject/service/plant"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	db, _ := config.ConnectDatabase()
	config.MigrateDB(db)

	e := echo.New()

	authJwt := middleware.JwtAlta{}
	authRepo := authRepo.NewAuthRepo(db)
	authService := authService.NewAuthService(authRepo, authJwt)
	authController := authController.NewAuthController(authService)

	plantRepo := plantRepo.NewPlantRepo(db)
	plantService := plantService.NewAuthService(plantRepo)
	plantController := plantController.NewPlantController(plantService)

	routeController := route.RouteController{
		AuthController:  *authController,
		PlantController: *plantController,
	}
	routeController.InitRoute(e)

	e.Start(":8000")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("failed load env")
	}
}
