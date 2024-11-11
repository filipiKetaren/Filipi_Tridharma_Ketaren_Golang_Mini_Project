package plantCondition

import (
	"miniproject/entities"
)

type PlantConditionRepoInterface interface {
	FindCondition(userID int) ([]entities.PlantCondition, error)
	FindConditionByID(plant entities.PlantCondition, conditionID, userID int) (entities.PlantCondition, error)
	CreateCondition(plant entities.PlantCondition) (entities.PlantCondition, error)
	UpdateCondition(plant entities.PlantCondition) (entities.PlantCondition, error)
	DeleteCondition(plant entities.PlantCondition) (entities.PlantCondition, error)
	FindPlantByIDAndUser(plantID, userID int) (entities.Plant, error)
	CheckPlantId(PlantID, userID int) error
}
