package plantCondition

import (
	"miniproject/controller/base"
	"miniproject/controller/plant_condition/request"
	"miniproject/controller/plant_condition/response"
	"miniproject/helper"
	pc "miniproject/service/plant_condition"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewPlantConditionController(pi pc.PlantConditionInterface) *PlantConditionController {
	return &PlantConditionController{
		plantConditionService: pi,
	}
}

type PlantConditionController struct {
	plantConditionService pc.PlantConditionInterface
}

func (plantConditionController PlantConditionController) FindController(c echo.Context) error {
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User ID not found in context"})
	}

	plantData, err := plantConditionController.plantConditionService.FindCondition(userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": plantData,
	})
}

func (plantConditionController PlantConditionController) FindByIdController(c echo.Context) error {
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User ID not found in context"})
	}

	conditionID, err := helper.GetIDParam(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	condition := request.PlantCondition{}
	plantData, err := plantConditionController.plantConditionService.FindConditionByID(condition.ToEntities(), conditionID, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.FromEntities(plantData))
}

func (plantConditionController PlantConditionController) CreateController(c echo.Context) error {
	// Tambahkan pengecekan awal untuk memastikan plantConditionService tidak nil
	if plantConditionController.plantConditionService == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Service is not initialized"})
	}

	condition := request.PlantCondition{}
	if err := c.Bind(&condition); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse request"})
	}

	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User ID tidak ditemukan dalam konteks"})
	}

	err := plantConditionController.plantConditionService.CheckPlantId(condition.PlantID, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	plantData, err := plantConditionController.plantConditionService.CreateCondition(condition.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.FromEntities(plantData))
}

func (plantConditionController PlantConditionController) UpdateController(c echo.Context) error {
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User ID not found in context"})
	}

	conditionID, err := helper.GetIDParam(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	conditionRequest := request.PlantCondition{}
	err = c.Bind(&conditionRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse request"})
	}

	err = plantConditionController.plantConditionService.CheckPlantId(conditionRequest.PlantID, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	condition := conditionRequest.ToEntities()
	condition.ID = conditionID
	plant, err := plantConditionController.plantConditionService.UpdateCondition(condition)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update plant"})
	}

	return base.SuccessResponse(c, response.FromEntities(plant))
}

func (plantConditionController PlantConditionController) DeleteController(c echo.Context) error {
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User ID not found in context"})
	}

	conditionID, err := helper.GetIDParam(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	condition := request.PlantCondition{}
	plantData, err := plantConditionController.plantConditionService.FindConditionByID(condition.ToEntities(), conditionID, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	err = c.Bind(&plantData)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	plant, err := plantConditionController.plantConditionService.DeleteCondition(plantData)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, response.FromEntities(plant))
}
