package response

import (
	"miniproject/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PlantCondition struct {
	ID               int       `json:"id"`
	PlantID          int       `json:"plant_id"`
	Plant            PlantData `json:"plant"`
	MoistureLevel    float32   `json:"moisture_level"`
	SunlightExposure string    `json:"sunlight_exposure"`
	Temperature      float32   `json:"temperature"`
	Notes            string    `json:"notes"`
}

type PlantConditionsResponse struct {
	Status    bool        `json:"status"`
	Message   string      `json:"message"`
	Condition []Condition `json:"data_plant_condition"`
}

type Condition struct {
	PlantCondition PlantCondition `json:"plant_condition"`
}

type PlantData struct {
	ID        int    `json:"id"`
	User      User   `json:"user"`
	PlantName string `json:"plant_name"`
	Species   string `json:"species"`
	Location  string `json:"location"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type PlantConditionSingularResponse struct {
	Status    bool                   `json:"status"`
	Message   string                 `json:"message"`
	Condition PlantConditionResponse `json:"data_plant_condition"`
}

type PlantConditionResponse struct {
	PlantCondition PlantCondition `json:"plant_condition"`
}

func FromEntities(plantCondition entities.PlantCondition) PlantCondition {
	return PlantCondition{
		ID:               plantCondition.ID,
		PlantID:          plantCondition.PlantID,
		MoistureLevel:    plantCondition.MoistureLevel,
		SunlightExposure: plantCondition.SunlightExposure,
		Temperature:      plantCondition.Temperature,
		Notes:            plantCondition.Notes,
	}
}

func SplitSliceResponse(plantData []entities.PlantCondition, plantEntities map[int]entities.Plant) []Condition {
	conditions := make([]Condition, 0)
	for _, condition := range plantData {
		plant := plantEntities[condition.PlantID]
		conditionData := Condition{
			PlantCondition: PlantCondition{
				ID:      condition.ID,
				PlantID: condition.PlantID,
				Plant: PlantData{
					ID: plant.ID,
					User: User{
						ID:       plant.UserID,
						Username: plant.User.Username, // Replace with actual username
						Email:    plant.User.Email,    // Replace with actual email
					},
					PlantName: plant.PlantName,
					Species:   plant.Species,
					Location:  plant.Location,
				},
				MoistureLevel:    condition.MoistureLevel,
				SunlightExposure: condition.SunlightExposure,
				Temperature:      condition.Temperature,
				Notes:            condition.Notes,
			},
		}
		conditions = append(conditions, conditionData)
	}
	return conditions
}

func SplitResponse(plantData entities.PlantCondition, plantEntity entities.Plant) PlantConditionResponse {
	plantConditionResponse := PlantConditionResponse{
		PlantCondition: PlantCondition{
			ID:      plantData.ID,
			PlantID: plantData.PlantID,
			Plant: PlantData{
				ID: plantEntity.ID,
				User: User{
					ID:       plantEntity.UserID,
					Username: plantEntity.User.Username, // Menggunakan username dari entitas Plant
					Email:    plantEntity.User.Email,    // Replace with actual email
				},
				PlantName: plantEntity.PlantName,
				Species:   plantEntity.Species,
				Location:  plantEntity.Location,
			},
			MoistureLevel:    plantData.MoistureLevel,
			SunlightExposure: plantData.SunlightExposure,
			Temperature:      plantData.Temperature,
			Notes:            plantData.Notes,
		},
	}
	return plantConditionResponse
}

func SuccessResponseCondition(c echo.Context, plantData PlantConditionResponse) error {
	return c.JSON(http.StatusOK, PlantConditionSingularResponse{
		Status:    true,
		Message:   "sukses",
		Condition: plantData,
	})
}
