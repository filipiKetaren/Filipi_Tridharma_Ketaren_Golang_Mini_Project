package suggestion

import "miniproject/entities"

type SuggestionRepoInterface interface {
	SaveSuggestion(suggestion entities.CareSuggestion) error
	CheckPlantExists(plantID int) (bool, error)
	FindSuggestion(userID int) ([]entities.CareSuggestion, error)
	GetAll(plants *[]entities.Plant) error
	GetPlantByID(plantID int) (entities.Plant, error)
}
