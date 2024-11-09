package plant

import (
	"miniproject/entities"
	"miniproject/repo/auth"
)

type Plant struct {
	ID        int       `gorm: primarykey`
	UserID    int       `gorm:"not null"`
	User      auth.User `gorm:"constraint:OnDelete:CASCADE;"`
	PlantName string    `gorm:"not null"`
	Species   string    `gorm:"not null"`
	Location  string    `gorm:"not null"`

	Plants []Plant `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

func FromEntities(plant entities.Plant) Plant {
	return Plant{
		ID:        plant.ID,
		UserID:    plant.UserID,
		PlantName: plant.PlantName,
		Species:   plant.Species,
		Location:  plant.Location,
	}
}

func (plant Plant) ToEntities() entities.Plant {
	return entities.Plant{
		ID:        plant.ID,
		UserID:    plant.UserID,
		PlantName: plant.PlantName,
		Species:   plant.Species,
		Location:  plant.Location,
	}
}
