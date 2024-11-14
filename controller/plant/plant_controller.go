package plant

import (
	"miniproject/controller/base"
	"miniproject/controller/plant/request"
	"miniproject/controller/plant/response"
	"miniproject/helper"
	"miniproject/service/auth"
	"miniproject/service/plant"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewPlantController(pi plant.PlantInterface, us auth.AuthInterface) *PlantController {
	return &PlantController{
		plantService: pi,
		userService:  us,
	}
}

type PlantController struct {
	plantService plant.PlantInterface
	userService  auth.AuthInterface
}

func (plantController PlantController) FindController(c echo.Context) error {
	userID := c.Get("user_id").(int) // Mendapatkan user_id dari context

	// Mengambil data plant dan user berdasarkan userID
	plantData, err := plantController.plantService.FindPlant(userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Mendapatkan data user yang sesuai dengan userID
	userData, err := plantController.userService.FindUserByIDs(userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Menggabungkan data plant dan user
	responseData := response.SplitSliceResponse(plantData, userData)

	// Mengembalikan response sukses dengan data yang sudah digabungkan
	return response.SliceSuccessResponse(c, responseData)
}

func (plantController PlantController) FindByIdController(c echo.Context) error {
	id, err := helper.GetIDParam(c)
	userID := c.Get("user_id").(int)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Ambil data plant berdasarkan ID dan userID
	plant, err := plantController.plantService.FindByIdPlant(id, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Jika plant ditemukan, ambil data user terkait
	user, err := plantController.userService.FindByUserID(plant.UserID) // Pastikan user di-fetch dengan benar
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	plantResponse := response.FromEntities(plant, user)

	// Kembalikan respons dengan data plant yang sudah dilengkapi dengan user
	return base.SuccessResponse(c, plantResponse)
}

func (plantController PlantController) CreateController(c echo.Context) error {
	plant := request.Plant{}
	userID := c.Get("user_id").(int)

	if err := c.Bind(&plant); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Set user_id pada data plant
	plant.UserID = userID

	// Panggil service untuk membuat data plant baru
	plantData, err := plantController.plantService.CreatePlant(plant.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Ambil data user terkait setelah plant berhasil dibuat
	user, err := plantController.userService.FindByUserID(plantData.UserID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Gabungkan data plant dengan data user dalam response
	plantResponse := response.FromEntities(plantData, user)

	// Mengembalikan respons sukses dengan data yang telah dibuat dan dilengkapi dengan user
	return base.SuccessResponse(c, plantResponse)
}

func (plantController PlantController) UpdateController(c echo.Context) error {
	id, err := helper.GetIDParam(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	userID := c.Get("user_id").(int)

	// Cek apakah user yang login dapat mengubah plant dengan ID tersebut
	plant, err := plantController.plantService.CheckUserLogin(id, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Bind data yang diterima
	plantData := request.FromEntities(plant)
	err = c.Bind(&plantData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse request"})
	}

	// Update plant dengan data yang telah di-bind
	plant, err = plantController.plantService.UpdatePlant(plantData.ToEntities())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update plant"})
	}

	// Ambil data user terkait setelah plant berhasil diperbarui
	user, err := plantController.userService.FindByUserID(plant.UserID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Gabungkan data plant dengan data user dalam response
	plantResponse := response.FromEntities(plant, user)

	// Mengembalikan respons sukses dengan data yang telah di-update dan dilengkapi dengan user
	return base.SuccessResponse(c, plantResponse)
}

func (plantController PlantController) DeleteController(c echo.Context) error {
	id, err := helper.GetIDParam(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	userID := c.Get("user_id").(int)

	// Cek apakah user yang login dapat menghapus plant dengan ID tersebut
	plant, err := plantController.plantService.CheckUserLogin(id, userID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Bind data plant yang akan dihapus
	err = c.Bind(&plant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse request"})
	}

	// Panggil service untuk menghapus plant
	plant, err = plantController.plantService.DeletePlant(plant)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete plant"})
	}

	// Ambil data user terkait setelah plant berhasil dihapus
	user, err := plantController.userService.FindByUserID(plant.UserID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	// Gabungkan data plant yang dihapus dengan data user dalam response
	plantResponse := response.FromEntities(plant, user)

	// Mengembalikan respons sukses setelah plant dihapus dan dilengkapi dengan user
	return base.SuccessResponse(c, plantResponse)
}
