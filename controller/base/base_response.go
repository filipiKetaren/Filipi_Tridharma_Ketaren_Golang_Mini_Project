package base

import (
	rp "miniproject/controller/plant_condition/response"
	"miniproject/controller/suggestion/response"
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
	Status    bool                      `json:"status"`
	Message   string                    `json:"message"`
	Condition rp.PlantConditionResponse `json:"data_plant_condition"`
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

func SuccessResponseSlice(c echo.Context, plantData []rp.PlantCondition) error {
	return c.JSON(http.StatusOK, rp.PlantConditionMajemukResponse{
		Status:    true,
		Message:   "sukses",
		Condition: plantData,
	})
}

func SuccessResponseSuggestion(c echo.Context, Suggestion response.CareSuggestionResponse) error {
	// Membungkus data ke dalam struct PlantResponses
	plantResponses := response.SuggestionResponse{
		Status:  true,
		Message: "sukses",
		Data:    Suggestion,
	}
	return c.JSON(http.StatusOK, plantResponses)
}

func SliceSuccessResponseSuggetion(c echo.Context, Suggestion []response.CareSuggestionResponse) error {
	// Membungkus data ke dalam struct PlantResponses
	plantResponses := response.SuggestionResponses{
		Status:  true,
		Message: "sukses",
		Data:    Suggestion,
	}
	return c.JSON(http.StatusOK, plantResponses)
}
