package base

import (
	"miniproject/controller/plant_condition/response"
	"miniproject/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PlantConditionsResponse struct {
	Status    bool                            `json:"status"`
	Message   string                          `json:"message"`
	Condition response.PlantConditionResponse `json:"data_plant_condition"`
}
type SliceConditionResponse struct {
	Status    bool        `json:"status"`
	Message   string      `json:"message"`
	Condition interface{} `json:"condition"`
	Data      interface{} `json:"data"`
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

func SliceSuccessResponse(c echo.Context, Condition []response.Condition) error {
	return c.JSON(http.StatusOK, response.PlantConditionsResponse{
		Status:    true,
		Message:   "sukses",
		Condition: Condition,
	})
}
