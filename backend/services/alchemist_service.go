package services

import (
	"alchemy-system/database"
    "alchemy-system/models"
	"errors"
	"gorm.io/gorm"
)

// CreateAlchemist adds a new alchemist to the database.
func CreateAlchemist(alchemist models.Alchemist) (models.Alchemist, error) {
	result := database.DB.Create(&alchemist)
	if result.Error != nil {
		return alchemist, result.Error
	}
	return alchemist, nil
}

// GetAllAlchemists returns all registered alchemists.
func GetAllAlchemists() ([]models.Alchemist, error) {
	var alchemists []models.Alchemist
	result := database.DB.Find(&alchemists)
	return alchemists, result.Error
}

// GetAlchemistByID returns one alchemist by ID.
func GetAlchemistByID(id uint) (models.Alchemist, error) {
	var alchemist models.Alchemist
	result := database.DB.First(&alchemist, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return alchemist, errors.New("alchemist not found")
	}

	return alchemist, result.Error
}

// UpdateAlchemist updates an existing alchemist’s data.
func UpdateAlchemist(id uint, updated models.Alchemist) (models.Alchemist, error) {
	var alchemist models.Alchemist
	if err := database.DB.First(&alchemist, id).Error; err != nil {
		return alchemist, errors.New("alchemist not found")
	}

	alchemist.Name = updated.Name
	alchemist.Rank = updated.Rank
	alchemist.Specialty = updated.Specialty

	if err := database.DB.Save(&alchemist).Error; err != nil {
		return alchemist, err
	}
	return alchemist, nil
}

// DeleteAlchemist removes an alchemist by ID.
func DeleteAlchemist(id uint) error {
	var alchemist models.Alchemist
	if err := database.DB.First(&alchemist, id).Error; err != nil {
		return errors.New("alchemist not found")
	}
	return database.DB.Delete(&alchemist).Error
}
