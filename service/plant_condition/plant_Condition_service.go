package plantCondition

import (
	"miniproject/entities"
	plantCondition "miniproject/repo/plant_condition"
)

func NewPlantConditionService(pr plantCondition.PlantConditionRepoInterface) *PlantConditionService {
	return &PlantConditionService{
		plantConditionRepoInterface: pr,
	}
}

type PlantConditionService struct {
	plantConditionRepoInterface plantCondition.PlantConditionRepoInterface
}

func (plantConditionService PlantConditionService) FindCondition(userID int) ([]entities.PlantCondition, error) {
	plant, err := plantConditionService.plantConditionRepoInterface.FindCondition(userID)
	if err != nil {
		return []entities.PlantCondition{}, err
	}
	return plant, nil
}

func (plantConditionService PlantConditionService) FindConditionByID(plant entities.PlantCondition, conditionID, userID int) (entities.PlantCondition, error) {
	condition, err := plantConditionService.plantConditionRepoInterface.FindConditionByID(plant, conditionID, userID)
	if err != nil {
		return entities.PlantCondition{}, err
	}
	return condition, nil
}

func (plantConditionService PlantConditionService) CreateCondition(condition entities.PlantCondition) (entities.PlantCondition, error) {
	plant, err := plantConditionService.plantConditionRepoInterface.CreateCondition(condition)
	if err != nil {
		return entities.PlantCondition{}, err
	}
	return plant, nil
}

func (plantConditionService PlantConditionService) UpdateCondition(plant entities.PlantCondition) (entities.PlantCondition, error) {
	plant, err := plantConditionService.plantConditionRepoInterface.UpdateCondition(plant)
	if err != nil {
		return entities.PlantCondition{}, err
	}
	return plant, nil
}

func (plantConditionService PlantConditionService) DeleteCondition(plant entities.PlantCondition) (entities.PlantCondition, error) {
	plant, err := plantConditionService.plantConditionRepoInterface.DeleteCondition(plant)
	if err != nil {
		return entities.PlantCondition{}, err
	}
	return plant, nil
}

func (plantConditionService PlantConditionService) GetPlantByIDAndUser(plantID, userID int) (entities.Plant, error) {
	plant, err := plantConditionService.plantConditionRepoInterface.FindPlantByIDAndUser(plantID, userID)
	if err != nil {
		return entities.Plant{}, err
	}
	return plant, nil
}

func (plantConditionService PlantConditionService) CheckPlantId(plantID, userID int) error {
	err := plantConditionService.plantConditionRepoInterface.CheckPlantId(plantID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (plantConditionService PlantConditionService) FindByID(plantID int) (entities.PlantCondition, error) {
	condition, err := plantConditionService.plantConditionRepoInterface.FindByID(plantID)
	if err != nil {
		return entities.PlantCondition{}, err
	}
	return condition, nil
}
