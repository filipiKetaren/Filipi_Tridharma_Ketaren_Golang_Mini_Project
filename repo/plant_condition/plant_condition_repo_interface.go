package plantCondition

import (
	"miniproject/entities"
)

type PlantConditionRepoInterface interface {
	FindCondition(userID int) ([]entities.PlantCondition, error)
	FindConditionByID(conditionID, userID int) (entities.PlantCondition, error)
	CreateCondition(plant entities.PlantCondition) (entities.PlantCondition, error)
	UpdateCondition(plant entities.PlantCondition) (entities.PlantCondition, error)
	DeleteCondition(plant entities.PlantCondition) error
	FindPlantByIDAndUser(plantID, userID int) (entities.Plant, error)
	CheckPlantId(PlantID, userID int) error
	FindByID(PlantID int) (entities.PlantCondition, error)
	FindPlantByID(plantID, userID int) (entities.Plant, error)
}
