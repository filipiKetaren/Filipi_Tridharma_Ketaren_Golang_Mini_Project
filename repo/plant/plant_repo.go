package plant

import (
	"miniproject/entities"

	"gorm.io/gorm"
)

func NewPlantRepo(db *gorm.DB) *PlantRepo {
	return &PlantRepo{
		db: db,
	}
}

type PlantRepo struct {
	db *gorm.DB
}

func (plantRepo PlantRepo) FindPlant(userID int) ([]entities.Plant, error) {
	var plantDB []entities.Plant
	result := plantRepo.db.Where("user_id = ?", userID).Find(&plantDB)
	if result.Error != nil {
		return []entities.Plant{}, result.Error
	}
	return plantDB, nil
}

func (plantRepo PlantRepo) FindByIdPlant(id int, userID int) (entities.Plant, error) {
	return plantRepo.CheckUserLogin(id, userID)
}

func (plantRepo PlantRepo) CreatePlant(plant entities.Plant) (entities.Plant, error) {
	plantDB := FromEntities(plant)
	result := plantRepo.db.Create(&plantDB)
	if result.Error != nil {
		return entities.Plant{}, result.Error
	}
	return plantDB.ToEntities(), nil
}

func (plantRepo PlantRepo) UpdatePlant(plant entities.Plant) (entities.Plant, error) {
	plantDB := FromEntities(plant)
	result := plantRepo.db.Updates(&plantDB)
	if result.Error != nil {
		return entities.Plant{}, result.Error
	}
	return plantDB.ToEntities(), nil
}

func (plantRepo PlantRepo) DeletePlant(plant entities.Plant) (entities.Plant, error) {
	plantDB := FromEntities(plant)
	result := plantRepo.db.Delete(&plantDB)
	if result.Error != nil {
		return entities.Plant{}, result.Error
	}
	return plantDB.ToEntities(), nil
}

func (plantRepo PlantRepo) CheckUserLogin(id, userID int) (entities.Plant, error) {
	var plantDB entities.Plant
	err := plantRepo.db.Where("id = ? AND user_id = ?", id, userID).First(&plantDB)
	if err.Error != nil {
		return entities.Plant{}, err.Error
	}
	return plantDB, nil
}
