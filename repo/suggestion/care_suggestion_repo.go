package suggestion

import (
	"fmt"
	"miniproject/entities"

	"gorm.io/gorm"
)

func NewCareSuggestionRepository(db *gorm.DB) *careSuggestionRepository {
	return &careSuggestionRepository{
		db: db}
}

type careSuggestionRepository struct {
	db *gorm.DB
}

func (repo careSuggestionRepository) SaveSuggestion(suggestion entities.CareSuggestion) error {
	if err := repo.db.Create(&suggestion).Error; err != nil {
		return fmt.Errorf("failed to save care suggestion: %w", err)
	}
	return nil
}

func (repo careSuggestionRepository) CheckPlantExists(plantID int) (bool, error) {
	var plant entities.Plant
	err := repo.db.First(&plant, plantID).Error
	if err != nil {
		return false, err // Terjadi error lainnya
	}
	return true, nil // Tanaman ditemukan
}

func (repo careSuggestionRepository) FindSuggestion(userID int) ([]entities.CareSuggestion, error) {
	var suggestions []entities.CareSuggestion
	result := repo.db.
		Joins("JOIN plants ON plants.id = care_suggestions.plant_id").
		Where("plants.user_id = ?", userID).
		Preload("Plant.User").
		Preload("Plant").
		Find(&suggestions)

	// Log untuk memeriksa hasil query
	if result.Error != nil {
		return nil, result.Error
	}
	// fmt.Println("Suggestions Data:", suggestions) // Menampilkan data suggestions yang dimuat
	return suggestions, nil
}

func (repo careSuggestionRepository) GetAll(plants *[]entities.Plant) error {
	// Mengambil data dari database menggunakan GORM
	err := repo.db.Preload("User").Find(plants).Error
	if err != nil {
		return err
	}
	return nil
}

func (s careSuggestionRepository) GetPlantByID(plantID int) (entities.Plant, error) {
	// Membuat objek untuk menyimpan data tanaman yang diambil
	var plant entities.Plant

	// Query untuk mengambil data tanaman berdasarkan plantID
	err := s.db.Preload("User").First(&plant, plantID).Error
	if err != nil {
		// Jika terjadi error (misalnya data tidak ditemukan)
		return plant, fmt.Errorf("failed to find plant with id %d: %w", plantID, err)
	}

	// Mengembalikan objek tanaman yang ditemukan
	return plant, nil
}
