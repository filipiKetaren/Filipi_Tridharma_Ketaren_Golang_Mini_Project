package response

import (
	"miniproject/entities"
)

type PlantCondition struct {
	ID               int     `json:"id"`
	PlantID          int     `json:"plant_id"`
	Date             string  `json:"date"` // Ubah ke string untuk sementara
	MoistureLevel    float32 `json:"moisture_level"`
	SunlightExposure string  `json:"sunlight_exposure"`
	Temperature      float32 `json:"temperature"`
	Notes            string  `json:"notes"`
}

func FromEntities(plantCondition entities.PlantCondition) PlantCondition {
	return PlantCondition{
		ID:               plantCondition.ID,
		PlantID:          plantCondition.PlantID,
		Date:             plantCondition.Date,
		MoistureLevel:    plantCondition.MoistureLevel,
		SunlightExposure: plantCondition.SunlightExposure,
		Temperature:      plantCondition.Temperature,
		Notes:            plantCondition.Notes,
	}
}
