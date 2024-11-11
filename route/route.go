package route

import (
	"miniproject/controller/auth"
	"miniproject/controller/plant"
	plantCondition "miniproject/controller/plant_condition"
	"miniproject/middleware"
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController           auth.AuthController
	PlantController          plant.PlantController
	PlantConditionController plantCondition.PlantConditionController
	Jwt                      middleware.JwtAlta
}

func (rc RouteController) InitRoute(e *echo.Echo) {
	e.POST("/login", rc.AuthController.LoginController)
	e.POST("/register", rc.AuthController.RegisterController)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	ePlant := eJwt.Group("/plants")
	ePlant.Use(rc.Jwt.GetUserID)
	ePlant.GET("", rc.PlantController.FindController)
	ePlant.GET("/:id", rc.PlantController.FindByIdController)
	ePlant.POST("", rc.PlantController.CreateController)
	ePlant.PUT("/:id", rc.PlantController.UpdateController)
	ePlant.DELETE("/:id", rc.PlantController.DeleteController)

	eCondition := eJwt.Group("/condition")
	eCondition.Use(rc.Jwt.GetUserID)
	eCondition.GET("", rc.PlantConditionController.FindController)
	eCondition.GET("/:id", rc.PlantConditionController.FindByIdController)
	eCondition.POST("", rc.PlantConditionController.CreateController)
	eCondition.PUT("/:id", rc.PlantConditionController.UpdateController)
	eCondition.DELETE("/:id", rc.PlantConditionController.DeleteController)
}
