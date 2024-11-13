package suggestion

import "miniproject/entities"

type SuggestionServiceInterface interface {
	FindSuggestion(userID int) ([]entities.CareSuggestion, error)
	SaveCareSuggestion(plantID int, suggestion string) error
	CheckPlantExists(plantID int) (bool, error)
}