package suggestion

import (
	"fmt"
	"miniproject/entities"

	"gorm.io/gorm"
)

func NewCareSuggestionRepository(db *gorm.DB) *careSuggestionRepository {
	return &careSuggestionRepository{
		db: db}
}

type careSuggestionRepository struct {
	db *gorm.DB
}

func (repo careSuggestionRepository) SaveSuggestion(suggestion entities.CareSuggestion) error {
	if err := repo.db.Create(&suggestion).Error; err != nil {
		return fmt.Errorf("failed to save care suggestion: %w", err)
	}
	return nil
}

func (repo careSuggestionRepository) CheckPlantExists(plantID int) (bool, error) {
	var plant entities.Plant
	err := repo.db.First(&plant, plantID).Error
	if err != nil {
		return false, err // Terjadi error lainnya
	}
	return true, nil // Tanaman ditemukan
}

func (repo careSuggestionRepository) FindSuggestion(userID int) ([]entities.CareSuggestion, error) {
	var Suggestions []entities.CareSuggestion
	result := repo.db.Joins("JOIN plants ON plants.id = care_suggestions.plant_id").
		Where("plants.user_id = ?", userID).Find(&Suggestions)
	if result.Error != nil {
		return []entities.CareSuggestion{}, result.Error
	}
	return Suggestions, nil
}
