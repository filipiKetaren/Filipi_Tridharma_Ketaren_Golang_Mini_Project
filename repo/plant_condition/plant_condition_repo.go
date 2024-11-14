package plantCondition

import (
	"errors"
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

	// Mengambil data terkait Plant dan User
	for i, condition := range conditions {
		var plant entities.Plant
		if err := plantConditionRepo.db.Preload("User").Where("id = ?", condition.PlantID).First(&plant).Error; err != nil {
			return nil, err
		}
		conditions[i].Plant = plant
	}

	return conditions, nil
}

func (plantConditionRepo PlantConditionRepo) FindConditionByID(conditionID, userID int) (entities.PlantCondition, error) {
	var condition entities.PlantCondition
	result := plantConditionRepo.db.
		Joins("JOIN plants ON plants.id = plant_conditions.plant_id").
		Joins("JOIN users ON users.id = plants.user_id").
		Where("plant_conditions.id = ? AND plants.user_id = ?", conditionID, userID).
		Preload("Plant.User"). // Preload relasi Plant dan User
		First(&condition)
	if result.Error != nil {
		return entities.PlantCondition{}, result.Error
	}
	return condition, nil
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

func (plantConditionRepo PlantConditionRepo) DeleteCondition(plant entities.PlantCondition) error {
	plantDB := FromEntities(plant)
	result := plantConditionRepo.db.Delete(&plantDB)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo PlantConditionRepo) FindPlantByIDAndUser(plantID, userID int) (entities.Plant, error) {
	var plant entities.Plant
	// Gunakan Preload("User") untuk memuat data User terkait
	if err := repo.db.Where("id = ? AND user_id = ?", plantID, userID).
		Preload("User").
		First(&plant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return plant, nil
		}
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

func (repo PlantConditionRepo) FindByID(plantID int) (entities.PlantCondition, error) {
	var condition entities.PlantCondition
	if err := repo.db.
		Preload("Plant.User"). // Preload User for the associated Plant
		Where("plant_id = ?", plantID).
		First(&condition).Error; err != nil {
		return entities.PlantCondition{}, err
	}
	return condition, nil
}

func (repo PlantConditionRepo) FindPlantByID(plantID, userID int) (entities.Plant, error) {
	var plant entities.Plant
	// Tambahkan Preload untuk user
	err := repo.db.Preload("User").Where("id = ? AND user_id = ?", plantID, userID).First(&plant).Error
	if err != nil {
		return entities.Plant{}, err
	}
	return plant, nil
}
