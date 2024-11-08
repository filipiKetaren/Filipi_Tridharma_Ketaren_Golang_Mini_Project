package route

import (
	"miniproject/controller/auth"
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController auth.AuthController
}

func (rc RouteController) InitRoute(e *echo.Echo) {
	e.POST("/login", rc.AuthController.LoginController)
	e.POST("/register", rc.AuthController.RegisterController)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	eUser := eJwt.Group("/users")
	eUser.GET("", rc.AuthController.LoginController)

}
