package response

import (
	"miniproject/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Model utama untuk CareSuggestion response
type CareSuggestionResponse struct {
	Plant      PlantData `json:"plant"`
	Suggestion string    `json:"suggestion"`
}

type SuggestionResponses struct {
	Status  bool                     `json:"status"`
	Message string                   `json:"message"`
	Data    []CareSuggestionResponse `json:"data"`
}

type SuggestionResponse struct {
	Status  bool                   `json:"status"`
	Message string                 `json:"message"`
	Data    CareSuggestionResponse `json:"data"`
}

// type Suggestions struct {
// 	SuggestionResponse CareSuggestionResponse `json:"data"`
// }

// Model untuk data Plant dan User terkait
type PlantData struct {
	ID        int    `json:"id"`
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
	for _, suggestion := range suggestions {
		// Pastikan plantEntities memiliki entry dengan key PlantID
		plant := plantEntities[suggestion.PlantID]
		// Bentuk response untuk setiap CareSuggestion
		careSuggestionResponse := CareSuggestionResponse{
			Plant: PlantData{
				ID: plant.ID,
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
	return careSuggestionResponses
}

func SliceSuccessResponse(c echo.Context, Suggestion []CareSuggestionResponse) error {
	// Membungkus data ke dalam struct PlantResponses
	plantResponses := SuggestionResponses{
		Status:  true,
		Message: "sukses",
		Data:    Suggestion,
	}
	return c.JSON(http.StatusOK, plantResponses)
}

func SuccessResponseSuggestion(c echo.Context, Suggestion CareSuggestionResponse) error {
	// Membungkus data ke dalam struct PlantResponses
	plantResponses := SuggestionResponse{
		Status:  true,
		Message: "sukses",
		Data:    Suggestion,
	}
	return c.JSON(http.StatusOK, plantResponses)
}
