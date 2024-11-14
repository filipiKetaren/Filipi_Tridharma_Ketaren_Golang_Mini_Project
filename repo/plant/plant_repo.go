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
	var plants []entities.Plant
	result := plantRepo.db.Table("plants").
		Select("plants.id, plants.user_id, plants.plant_name, plants.species, plants.location, users.id as user_id, users.username, users.email, users.password").
		Joins("JOIN users ON users.id = plants.user_id").
		Find(&plants)

	if result.Error != nil {
		return nil, result.Error
	}
	return plants, nil
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

func (r PlantRepo) FindPlantsByUserID(userID int) ([]entities.Plant, error) {
	var plants []entities.Plant
	err := r.db.Where("user_id = ?", userID).Find(&plants).Error
	if err != nil {
		return nil, err
	}
	return plants, nil
}
