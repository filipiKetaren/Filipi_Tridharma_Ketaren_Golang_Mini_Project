package response

import (
	"miniproject/entities"
)

// Model utama untuk CareSuggestion response
type CareSuggestionResponse struct {
	ID         int       `json:"id"`
	PlantID    int       `json:"plant_id"`
	Plant      PlantData `json:"plant"`
	Suggestion string    `json:"suggestion"`
}

type Suggestions struct {
	SuggestionResponse CareSuggestionResponse `json:"Data"`
}

// Model untuk data Plant dan User terkait
type PlantData struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	User      User   `json:"user"`
	PlantName string `json:"plant_name"`
	Species   string `json:"species"`
	Location  string `json:"location"`
}

// Model untuk User
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func SplitSliceResponse(suggestions []entities.CareSuggestion, plantEntities map[int]entities.Plant) []CareSuggestionResponse {
	careSuggestionResponses := make([]CareSuggestionResponse, 0)

	// Log untuk memeriksa data suggestions yang diterima
	// fmt.Println("Suggestions received:", suggestions)

	for _, suggestion := range suggestions {
		// Pastikan plantEntities memiliki entry dengan key PlantID
		plant := plantEntities[suggestion.PlantID]
		// Bentuk response untuk setiap CareSuggestion
		careSuggestionResponse := CareSuggestionResponse{
			ID:      suggestion.ID,
			PlantID: suggestion.PlantID,
			Plant: PlantData{
				ID:     plant.ID,
				UserID: plant.UserID,
				User: User{
					ID:       plant.User.ID,
					Username: plant.User.Username,
					Email:    plant.User.Email,
				},
				PlantName: plant.PlantName,
				Species:   plant.Species,
				Location:  plant.Location,
			},
			Suggestion: suggestion.Suggestion,
		}
		careSuggestionResponses = append(careSuggestionResponses, careSuggestionResponse)
	}

	// // Log untuk memeriksa data response yang terbentuk
	// fmt.Println("CareSuggestion Responses:", careSuggestionResponses)

	return careSuggestionResponses
}
