package suggestion

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"miniproject/controller/base"
	"miniproject/controller/suggestion/response"
	"miniproject/entities"
	"miniproject/helper"
	plantCondition "miniproject/service/plant_condition"
	"miniproject/service/suggestion"
	"net/http"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

func NewSuggestionAIController(crs suggestion.SuggestionServiceInterface, pc plantCondition.PlantConditionInterface) *PlantSuggestionController {
	return &PlantSuggestionController{
		careSuggestionService: crs,
		plantConditionService: pc,
	}
}

type PlantSuggestionController struct {
	plantConditionService plantCondition.PlantConditionInterface
	careSuggestionService suggestion.SuggestionServiceInterface
}

func (controller PlantSuggestionController) GetCareSuggestion(c echo.Context) error {
	// Ambil `plant_id` dari parameter URL
	plantID, err := helper.GetIDParam(c)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "id plant condition not found"})
	}

	// Temukan kondisi tanaman berdasarkan `plant_id`
	plantCondition, err := controller.plantConditionService.FindByID(plantID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Plant condition not found"})
	}

	// Periksa apakah plant ID ada
	exists, err := controller.careSuggestionService.CheckPlantExists(plantID)
	if err != nil || !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Plant ID not found in plants table"})
	}

	// Siapkan konteks dan client AI
	ctx := context.Background()
	apiKey := os.Getenv("TOKEN_AI")
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal("Failed to create AI client:", err)
	}
	defer client.Close()

	// Tentukan model AI yang digunakan
	model := client.GenerativeModel("gemini-pro")
	model.SetTemperature(0.7)

	// Siapkan kondisi tanaman untuk dikirim ke model AI
	genAIParts := []genai.Part{
		genai.Text("Kelembapan: " + fmt.Sprintf("%.1f", plantCondition.MoistureLevel)),
		genai.Text("Paparan Sinar Matahari: " + plantCondition.SunlightExposure),
		genai.Text("Suhu: " + fmt.Sprintf("%.1f", plantCondition.Temperature)),
		genai.Text("Catatan: " + plantCondition.Notes),
	}

	// Tambahkan pertanyaan tentang rekomendasi perawatan
	question := genai.Text("Apa rekomendasi perawatan untuk kondisi tanaman di atas?")

	// Kirim permintaan ke model AI
	resp, err := model.GenerateContent(ctx, append(genAIParts, question)...)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get care suggestion"})
	}

	// Ambil rekomendasi perawatan dari respons AI
	careSuggestion := resp.Candidates[0].Content.Parts[0]

	// Jika careSuggestion adalah tipe data selain string (misalnya map atau struct), sesuaikan
	// Misalnya, jika careSuggestion adalah objek JSON atau map, Anda bisa menyimpan sebagai JSON string
	// Menggunakan json.Marshal untuk mengkonversi tipe data ke string
	careSuggestionJSON, err := json.Marshal(careSuggestion)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to marshal care suggestion"})
	}

	// Simpan saran perawatan ke dalam database
	err = controller.careSuggestionService.SaveCareSuggestion(plantID, string(careSuggestionJSON))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save care suggestion"})
	}

	// Ambil data tanaman dan user terkait
	plant, err := controller.careSuggestionService.GetPlantByID(plantID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get plant data"})
	}

	// Menyusun respons dengan data tanaman, user, dan saran perawatan
	response := response.CareSuggestionResponse{
		ID:      plantID,
		PlantID: plantID,
		Plant: response.PlantData{
			ID:     plant.ID,
			UserID: plant.UserID,
			User: response.User{
				ID:       plant.User.ID,
				Username: plant.User.Username,
				Email:    plant.User.Email,
			},
			PlantName: plant.PlantName,
			Species:   plant.Species,
			Location:  plant.Location,
		},
		Suggestion: string(careSuggestionJSON),
	}

	// Kirimkan hasil rekomendasi ke client
	return c.JSON(http.StatusOK, map[string]interface{}{"data": response})
}

func (plantSuggestionController PlantSuggestionController) FindController(c echo.Context) error {
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User ID not found in context"})
	}

	// Ambil data saran perawatan berdasarkan user ID
	suggestionData, err := plantSuggestionController.careSuggestionService.FindSuggestion(userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Log untuk memeriksa data suggestionData
	// fmt.Println("Suggestion Data:", suggestionData)

	// Ambil data plants dan preload data User
	var plants []entities.Plant
	err = plantSuggestionController.careSuggestionService.GetPlants(&plants)
	if err != nil {
		fmt.Println("Error fetching plants:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error fetching plant data"})
	}

	// Inisialisasi map plantEntities
	plantEntities := make(map[int]entities.Plant)
	for _, plant := range plants {
		plantEntities[plant.ID] = plant
	}

	// Panggil SplitSliceResponse untuk membentuk response yang sesuai
	responseData := response.SplitSliceResponse(suggestionData, plantEntities)

	// Log untuk memeriksa data responseData
	// fmt.Println("Response Data:", responseData)

	return c.JSON(http.StatusOK, map[string]interface{}{"data": responseData})
}
