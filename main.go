package main

import (
	"log"
	"miniproject/config"
	authController "miniproject/controller/auth"
	plantController "miniproject/controller/plant"
	conditionController "miniproject/controller/plant_condition"
	"miniproject/controller/suggestion"
	"miniproject/middleware"
	authRepo "miniproject/repo/auth"
	plantRepo "miniproject/repo/plant"
	conditionRepo "miniproject/repo/plant_condition"
	suggestionRepo "miniproject/repo/suggestion"
	"miniproject/route"
	authService "miniproject/service/auth"
	plantService "miniproject/service/plant"
	conditionService "miniproject/service/plant_condition"
	suggestionService "miniproject/service/suggestion"

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

	conditionRepo := conditionRepo.NewPlantConditionRepo(db)
	conditionService := conditionService.NewPlantConditionService(conditionRepo)
	conditionController := conditionController.NewPlantConditionController(conditionService)

	suggestionAIRepo := suggestionRepo.NewCareSuggestionRepository(db)
	suggestionAIService := suggestionService.NewSuggestionService(suggestionAIRepo)
	suggestionAIController := suggestion.NewSuggestionAIController(suggestionAIService, conditionService)

	routeController := route.RouteController{
		AuthController:           *authController,
		PlantController:          *plantController,
		PlantConditionController: *conditionController,
		SuggestionController:     *suggestionAIController,
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
