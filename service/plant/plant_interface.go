package plant

import "miniproject/entities"

type PlantInterface interface {
	FindPlant(userID int) ([]entities.Plant, error)
	FindByIdPlant(id int, userID int) (entities.Plant, error)
	CreatePlant(plant entities.Plant) (entities.Plant, error)
	UpdatePlant(plant entities.Plant) (entities.Plant, error)
	DeletePlant(plant entities.Plant) (entities.Plant, error)
	CheckUserLogin(id, userID int) (entities.Plant, error)
	// FindPlantsByUserID(userID int) ([]entities.User, error)
}
