package plant

import (
	"miniproject/entities"
)

type plant struct {
	ID        int    `gorm: "primarykey"`
	UserID    int    `gorm:"not null"`
	PlantName string `gorm:"not null"`
	Species   string `gorm:"not null"`
	Location  string `gorm:"not null"`
}

func FromEntities(plants entities.Plant) plant {
	return plant{
		ID:        plants.ID,
		UserID:    plants.UserID,
		PlantName: plants.PlantName,
		Species:   plants.Species,
		Location:  plants.Location,
	}
}

func (plant plant) ToEntities() entities.Plant {
	return entities.Plant{
		ID:        plant.ID,
		UserID:    plant.UserID,
		PlantName: plant.PlantName,
		Species:   plant.Species,
		Location:  plant.Location,
	}
}
