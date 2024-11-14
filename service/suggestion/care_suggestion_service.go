package suggestion

import (
	"fmt"
	"miniproject/entities"
	"miniproject/repo/suggestion"
)

func NewSuggestionService(crs suggestion.SuggestionRepoInterface) *CareSuggestionService {
	return &CareSuggestionService{
		suggestionRepo: crs}
}

type CareSuggestionService struct {
	suggestionRepo suggestion.SuggestionRepoInterface
}

// Fungsi untuk meminta saran perawatan ke API Gen AI
func (c CareSuggestionService) SaveCareSuggestion(plantID int, suggestion string) error {
	plantExists, err := c.suggestionRepo.CheckPlantExists(plantID)
	if err != nil || !plantExists {
		return fmt.Errorf("plant with id %d not found", plantID)
	}

	// Membuat objek CareSuggestion
	careSuggestion := entities.CareSuggestion{
		PlantID:    plantID,
		Suggestion: suggestion,
	}

	// Menyimpan objek CareSuggestion ke repository
	return c.suggestionRepo.SaveSuggestion(careSuggestion)
}

func (c CareSuggestionService) CheckPlantExists(plantID int) (bool, error) {
	return c.suggestionRepo.CheckPlantExists(plantID)
}

func (c CareSuggestionService) FindSuggestion(userID int) ([]entities.CareSuggestion, error) {
	suggestion, err := c.suggestionRepo.FindSuggestion(userID)
	if err != nil {
		return []entities.CareSuggestion{}, err
	}
	return suggestion, nil
}

func (c CareSuggestionService) GetPlants(plants *[]entities.Plant) error {
	// Mengambil data tanaman menggunakan repository
	err := c.suggestionRepo.GetAll(plants)
	if err != nil {
		return fmt.Errorf("error fetching plants: %w", err)
	}
	return nil
}

func (c CareSuggestionService) GetPlantByID(plantID int) (entities.Plant, error) {
	return c.suggestionRepo.GetPlantByID(plantID)
}
