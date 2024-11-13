package suggestion

import (
	"miniproject/entities"
)

type CareSuggestion struct {
	ID         int    `gorm:"primaryKey"`
	PlantID    int    `gorm:"not null"`
	Suggestion string `gorm:"type:text;not null"` // Menyimpan saran sebagai string
}

func FromEntities(careSuggestion entities.CareSuggestion) CareSuggestion {
	return CareSuggestion{
		ID:         careSuggestion.ID,
		PlantID:    careSuggestion.PlantID,
		Suggestion: careSuggestion.Suggestion,
	}
}

func (careSuggestion CareSuggestion) ToEntities() entities.CareSuggestion {
	return entities.CareSuggestion{
		ID:         careSuggestion.ID,
		PlantID:    careSuggestion.PlantID,
		Suggestion: careSuggestion.Suggestion,
	}
}
