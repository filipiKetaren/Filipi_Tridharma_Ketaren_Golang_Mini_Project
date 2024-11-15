package response

import (
	"miniproject/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Plant struct {
	ID        int
	User      User
	PlantName string
	Species   string
	Location  string
}

// Struktur data untuk User
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Struktur data untuk Plant dengan user terkait
type PlantWithUser struct {
	Plant struct {
		ID        int    `json:"id"`
		User      User   `json:"user"`
		PlantName string `json:"plant_name"`
		Species   string `json:"species"`
		Location  string `json:"location"`
	} `json:"plant"`
}

type PlantResponse struct {
	Plant Plant `json:"plant"`
}

type PlantResponses struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    []PlantResponse `json:"data_plant"`
}

// Fungsi untuk mengubah data Plant dari entities ke format response
func FromEntities(plant entities.Plant, user entities.User) PlantWithUser {
	return PlantWithUser{
		Plant: struct {
			ID        int    `json:"id"`
			User      User   `json:"user"`
			PlantName string `json:"plant_name"`
			Species   string `json:"species"`
			Location  string `json:"location"`
		}{
			ID: plant.ID,
			User: User{
				ID:       user.ID,
				Username: user.Username,
				Email:    user.Email,
			},
			PlantName: plant.PlantName,
			Species:   plant.Species,
			Location:  plant.Location,
		},
	}
}

func FromEntitiesNoPassword(user entities.User) User {
	return User{
		ID:       user.ID,
		Username: user.Password,
		Email:    user.Email,
	}
}

func SplitSliceResponse(plantData []entities.Plant, userData []entities.User) []PlantResponse {
	var plantResponses []PlantResponse
	for _, plant := range plantData {
		// Temukan user yang sesuai dengan userID
		var user User
		for _, u := range userData {
			if u.ID == plant.UserID {
				user = User{
					ID:       u.ID,
					Username: u.Username,
					Email:    u.Email,
				}
				break
			}
		}

		plantResponses = append(plantResponses, PlantResponse{
			Plant: Plant{
				ID:        plant.ID,
				PlantName: plant.PlantName,
				Species:   plant.Species,
				Location:  plant.Location,
				User:      user, // Menambahkan user yang sesuai
			},
		})
	}
	return plantResponses
}

func SliceSuccessResponse(c echo.Context, plantResponse []PlantResponse) error {
	// Membungkus data ke dalam struct PlantResponses
	plantResponses := PlantResponses{
		Status:  true,
		Message: "sukses",
		Data:    plantResponse,
	}
	return c.JSON(http.StatusOK, plantResponses)
}
