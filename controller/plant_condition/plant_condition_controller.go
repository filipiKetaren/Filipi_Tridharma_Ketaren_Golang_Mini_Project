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

	// Get plant condition data
	plantData, err := plantConditionController.plantConditionService.FindCondition(userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// // Debugging tambahan untuk memastikan data User dimuat
	// for _, condition := range plantData {
	// 	fmt.Printf("Condition ID: %d, Plant ID: %d, User ID: %d, Username: %s\n",
	// 		condition.ID, condition.Plant.ID, condition.Plant.User.ID, condition.Plant.User.Username)
	// }

	// Prepare data for response as per the new structure
	var result []response.Condition
	for _, condition := range plantData {
		// Pastikan User dimuat dengan benar
		result = append(result, response.Condition{
			PlantCondition: response.PlantCondition{
				ID:      condition.ID,
				PlantID: condition.PlantID,
				Plant: response.PlantData{
					ID: condition.Plant.ID,
					User: response.User{
						ID:       condition.Plant.User.ID,
						Username: condition.Plant.User.Username,
						Email:    condition.Plant.User.Email,
					},
					PlantName: condition.Plant.PlantName,
					Species:   condition.Plant.Species,
					Location:  condition.Plant.Location,
				},
				MoistureLevel:    condition.MoistureLevel,
				SunlightExposure: condition.SunlightExposure,
				Temperature:      condition.Temperature,
				Notes:            condition.Notes,
			},
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data_plant_condition": result,
	})
}

func (plantConditionController PlantConditionController) FindByIdController(c echo.Context) error {
	// Get user_id from context
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User ID not found in context"})
	}

	// Get conditionID from URL parameters
	conditionID, err := helper.GetIDParam(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Get the plant condition by ID

	plantData, err := plantConditionController.plantConditionService.FindConditionByID(conditionID, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Get associated plant entity for this condition
	plantEntity, err := plantConditionController.plantConditionService.FindPlantByID(plantData.PlantID, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Call SplitResponse with both the condition data and plant entity
	responseCondition := response.SplitResponse(plantData, plantEntity)

	return response.SuccessResponseCondition(c, responseCondition)
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

	// Cek jika PlantID valid untuk user ini
	err := plantConditionController.plantConditionService.CheckPlantId(condition.PlantID, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Create the plant condition
	plantData, err := plantConditionController.plantConditionService.CreateCondition(condition.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Fetch the associated plant data
	plantEntity, err := plantConditionController.plantConditionService.FindPlantByID(plantData.PlantID, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Format the response using SplitResponse
	responseCondition := response.SplitResponse(plantData, plantEntity)

	// Return the response
	return response.SuccessResponseCondition(c, responseCondition)
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

	// Validate PlantID for the current user
	err = plantConditionController.plantConditionService.CheckPlantId(conditionRequest.PlantID, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Prepare the condition entity for update
	condition := conditionRequest.ToEntities()
	condition.ID = conditionID

	// Update the condition
	updatedCondition, err := plantConditionController.plantConditionService.UpdateCondition(condition)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update plant condition"})
	}

	// Fetch the associated plant data
	plantEntity, err := plantConditionController.plantConditionService.FindPlantByID(updatedCondition.PlantID, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Format the response using SplitResponse
	responseCondition := response.SplitResponse(updatedCondition, plantEntity)

	// Return the response
	return response.SuccessResponseCondition(c, responseCondition)
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

	plantData, err := plantConditionController.plantConditionService.FindConditionByID(conditionID, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Delete the condition
	err = plantConditionController.plantConditionService.DeleteCondition(plantData)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Fetch the associated plant data after deletion
	plantEntity, err := plantConditionController.plantConditionService.FindPlantByID(plantData.PlantID, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Format the response using SplitResponse for the deleted condition
	responseCondition := response.SplitResponse(plantData, plantEntity)

	// Return the response
	return response.SuccessResponseCondition(c, responseCondition)
}
