package response

import "miniproject/entities"

type Plant struct {
	ID        int    `json: "id"`
	UserID    int    `json: "user_id"`
	PlantName string `json: "plant_name"`
	Species   string `json: "species"`
	Location  string `json: "location"`
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
