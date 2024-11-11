package plantCondition

import (
	"miniproject/entities"

	"gorm.io/gorm"
)

func NewPlantConditionRepo(db *gorm.DB) *PlantConditionRepo {
	return &PlantConditionRepo{
		db: db,
	}
}

type PlantConditionRepo struct {
	db *gorm.DB
}

func (plantConditionRepo PlantConditionRepo) FindCondition(userID int) ([]entities.PlantCondition, error) {
	var conditions []entities.PlantCondition
	result := plantConditionRepo.db.Joins("JOIN plants ON plants.id = plant_conditions.plant_id").
		Where("plants.user_id = ?", userID).Find(&conditions)
	if result.Error != nil {
		return []entities.PlantCondition{}, result.Error
	}
	return conditions, nil
}

func (plantConditionRepo PlantConditionRepo) FindConditionByID(condition entities.PlantCondition, conditionID, userID int) (entities.PlantCondition, error) {
	conditionDB := FromEntities(condition)
	result := plantConditionRepo.db.Joins("JOIN plants ON plants.id = plant_conditions.plant_id").
		Where("plant_conditions.id = ? AND plants.user_id = ?", conditionID, userID).
		First(&conditionDB)
	if result.Error != nil {
		return entities.PlantCondition{}, result.Error
	}
	return conditionDB.ToEntities(), nil
}

func (plantConditionRepo PlantConditionRepo) CreateCondition(condition entities.PlantCondition) (entities.PlantCondition, error) {
	plantDB := FromEntities(condition)
	result := plantConditionRepo.db.Create(&plantDB)
	if result.Error != nil {
		return entities.PlantCondition{}, result.Error
	}

	return plantDB.ToEntities(), nil
}

func (plantConditionRepo PlantConditionRepo) UpdateCondition(condition entities.PlantCondition) (entities.PlantCondition, error) {
	plantDB := FromEntities(condition)
	result := plantConditionRepo.db.Model(&plantDB).Where("id = ?", condition.ID).Updates(condition)
	if result.Error != nil {
		return entities.PlantCondition{}, result.Error
	}
	return plantDB.ToEntities(), nil
}

func (plantConditionRepo PlantConditionRepo) DeleteCondition(plant entities.PlantCondition) (entities.PlantCondition, error) {
	plantDB := FromEntities(plant)
	result := plantConditionRepo.db.Delete(&plantDB)
	if result.Error != nil {
		return entities.PlantCondition{}, result.Error
	}
	return plantDB.ToEntities(), nil
}

func (repo PlantConditionRepo) FindPlantByIDAndUser(plantID, userID int) (entities.Plant, error) {
	var plant entities.Plant
	if err := repo.db.Where("id = ? AND user_id = ?", plantID, userID).First(&plant).Error; err != nil {
		return plant, err
	}
	return plant, nil
}

func (repo PlantConditionRepo) CheckPlantId(plantID, userID int) error {
	var plant entities.Plant
	if err := repo.db.Where("id = ? AND user_id = ?", plantID, userID).First(&plant).Error; err != nil {
		return err
	}
	return nil
}
