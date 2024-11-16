package response

import (
	"miniproject/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PlantCondition struct {
	// ID               int       `json:"id"`
	// PlantID          int       `json:"plant_id"`
	Plant            PlantData `json:"plant"`
	MoistureLevel    float32   `json:"moisture_level"`
	SunlightExposure string    `json:"sunlight_exposure"`
	Temperature      float32   `json:"temperature"`
	Notes            string    `json:"notes"`
}

type PlantConditionsResponse struct {
	Status    bool        `json:"status"`
	Message   string      `json:"message"`
	Condition []Condition `json:"data"`
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
	Status    bool           `json:"status"`
	Message   string         `json:"message"`
	Condition PlantCondition `json:"data"`
}

type PlantConditionMajemukResponse struct {
	Status    bool             `json:"status"`
	Message   string           `json:"message"`
	Condition []PlantCondition `json:"data"`
}

type PlantConditionResponse struct {
	PlantCondition PlantCondition `json:"plant_condition"`
}

func FromEntities(plantCondition entities.PlantCondition) PlantCondition {
	return PlantCondition{
		MoistureLevel:    plantCondition.MoistureLevel,
		SunlightExposure: plantCondition.SunlightExposure,
		Temperature:      plantCondition.Temperature,
		Notes:            plantCondition.Notes,
	}
}

func SplitSliceResponse(plantData []entities.PlantCondition) []PlantCondition {
	var result []PlantCondition
	for _, condition := range plantData {
		// Pastikan User dimuat dengan benar
		result = append(result, PlantCondition{
			Plant: PlantData{
				ID: condition.Plant.ID,
				User: User{
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
		})
	}
	return result
}

func SplitResponse(plantData entities.PlantCondition, plantEntity entities.Plant) PlantCondition {
	PlantCondition := PlantCondition{
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
	}
	return PlantCondition
}

func SuccessResponseCondition(c echo.Context, plantData PlantCondition) error {
	return c.JSON(http.StatusOK, PlantConditionSingularResponse{
		Status:    true,
		Message:   "sukses",
		Condition: plantData,
	})
}

func SuccessResponseSlice(c echo.Context, plantData []PlantCondition) error {
	return c.JSON(http.StatusOK, PlantConditionMajemukResponse{
		Status:    true,
		Message:   "sukses",
		Condition: plantData,
	})
}
