package plant

import (
	"miniproject/controller/base"
	"miniproject/controller/plant/request"
	"miniproject/controller/plant/response"
	"miniproject/helper"
	"miniproject/service/plant"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewPlantController(pi plant.PlantInterface) *PlantController {
	return &PlantController{
		plantService: pi,
	}
}

type PlantController struct {
	plantService plant.PlantInterface
}

func (plantController PlantController) FindController(c echo.Context) error {
	plant := request.Plant{}
	userID := c.Get("user_id").(int)
	c.Bind(&plant)
	plantData, err := plantController.plantService.FindPlant(userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": plantData,
	})
}

func (plantController PlantController) FindByIdController(c echo.Context) error {
	id, err := helper.GetIDParam(c)
	userID := c.Get("user_id").(int)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	plant, err := plantController.plantService.FindByIdPlant(id, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, response.FromEntities(plant))
}

func (plantController PlantController) CreateController(c echo.Context) error {
	plant := request.Plant{}

	// Mendapatkan user_id dari context
	userID := c.Get("user_id").(int)
	if err := c.Bind(&plant); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Menetapkan user_id pada data plant
	plant.UserID = userID

	plantData, err := plantController.plantService.CreatePlant(plant.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccessResponse(c, response.FromEntities(plantData))
}

func (plantController PlantController) UpdateController(c echo.Context) error {
	id, err := helper.GetIDParam(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	userID := c.Get("user_id").(int)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	plant, err := plantController.plantService.CheckUserLogin(id, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	err = c.Bind(&plant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse request"})
	}

	plant, err = plantController.plantService.UpdatePlant(plant)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update plant"})
	}

	return base.SuccessResponse(c, response.FromEntities(plant))
}

func (plantController PlantController) DeleteController(c echo.Context) error {
	id, err := helper.GetIDParam(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	userID := c.Get("user_id").(int)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	plant, err := plantController.plantService.CheckUserLogin(id, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	err = c.Bind(&plant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse request"})
	}

	plant, err = plantController.plantService.DeletePlant(plant)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update plant"})
	}

	return base.SuccessResponse(c, response.FromEntities(plant))
}
