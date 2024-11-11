package plantCondition

import (
	"miniproject/entities"
)

type plantCondition struct {
	ID               int    `gorm:"primaryKey"`
	PlantID          int    `gorm:"not null"` // Foreign key to Plant
	Date             string `gorm:"not null"`
	MoistureLevel    float32
	SunlightExposure string
	Temperature      float32
	Notes            string
}

func FromEntities(plantConditions entities.PlantCondition) plantCondition {
	return plantCondition{
		ID:               plantConditions.ID,
		PlantID:          plantConditions.PlantID,
		Date:             plantConditions.Date,
		MoistureLevel:    plantConditions.MoistureLevel,
		SunlightExposure: plantConditions.SunlightExposure,
		Temperature:      plantConditions.Temperature,
		Notes:            plantConditions.Notes,
	}
}

func (plantCondition plantCondition) ToEntities() entities.PlantCondition {
	return entities.PlantCondition{
		ID:               plantCondition.ID,
		PlantID:          plantCondition.PlantID,
		Date:             plantCondition.Date,
		MoistureLevel:    plantCondition.MoistureLevel,
		SunlightExposure: plantCondition.SunlightExposure,
		Temperature:      plantCondition.Temperature,
		Notes:            plantCondition.Notes,
	}
}
