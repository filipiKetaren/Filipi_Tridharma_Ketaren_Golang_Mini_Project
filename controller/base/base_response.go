package base

import (
	"miniproject/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status  bool        `json: "status`
	Message string      `json: "message"`
	Data    interface{} `json: "data"`
}

func SuccessResponse(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "sukses",
		Data:    data,
	})
}

func ErrorResponse(c echo.Context, err error) error {
	return c.JSON(helper.GetResponseCodeFromErr(err), BaseResponse{
		Status:  false,
		Message: err.Error(),
	})
}
