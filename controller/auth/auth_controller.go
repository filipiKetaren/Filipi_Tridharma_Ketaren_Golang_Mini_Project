package auth

import (
	"miniproject/controller/auth/request"
	"miniproject/controller/auth/response"
	"miniproject/controller/base"
	"miniproject/service/auth"

	"github.com/labstack/echo/v4"
)

func NewAuthController(as auth.AuthInterface) *AuthController {
	return &AuthController{
		authService: as,
	}
}

type AuthController struct {
	authService auth.AuthInterface
}

func (userController AuthController) LoginController(c echo.Context) error {
	userLogin := request.LoginRequest{}
	c.Bind(&userLogin)
	user, err := userController.authService.Login(userLogin.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, response.FromEntities(user))
}
