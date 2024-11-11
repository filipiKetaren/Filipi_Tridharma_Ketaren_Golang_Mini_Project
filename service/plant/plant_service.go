package plant

import (
	"miniproject/entities"
	"miniproject/repo/plant"
)

func NewAuthService(pr plant.PlantRepoInterface) *PlantService {
	return &PlantService{
		plantRepoInterface: pr,
	}
}

type PlantService struct {
	plantRepoInterface plant.PlantRepoInterface
}

func (plantService PlantService) FindPlant(userID int) ([]entities.Plant, error) {
	plant, err := plantService.plantRepoInterface.FindPlant(userID)
	if err != nil {
		return []entities.Plant{}, err
	}
	return plant, nil
}

func (plantService PlantService) FindByIdPlant(id int, userID int) (entities.Plant, error) {
	return plantService.CheckUserLogin(id, userID)
}

func (plantService PlantService) CreatePlant(plant entities.Plant) (entities.Plant, error) {
	plant, err := plantService.plantRepoInterface.CreatePlant(plant)
	if err != nil {
		return entities.Plant{}, err
	}
	return plant, nil
}

func (plantService PlantService) UpdatePlant(plant entities.Plant) (entities.Plant, error) {
	plant, err := plantService.plantRepoInterface.UpdatePlant(plant)
	if err != nil {
		return entities.Plant{}, err
	}
	return plant, nil
}

func (plantService PlantService) DeletePlant(plant entities.Plant) (entities.Plant, error) {
	plant, err := plantService.plantRepoInterface.DeletePlant(plant)
	if err != nil {
		return entities.Plant{}, err
	}
	return plant, nil
}

func (plantService PlantService) CheckUserLogin(id, userID int) (entities.Plant, error) {
	plant, err := plantService.plantRepoInterface.CheckUserLogin(id, userID)
	if err != nil {
		return entities.Plant{}, err
	}
	return plant, nil
}
