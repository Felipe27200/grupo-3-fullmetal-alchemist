package services

import (
	"alchemy-system/backend/database"
	"alchemy-system/backend/models"
	"errors"

	"gorm.io/gorm"
)

// CreateMission inserts a new mission into the database
func CreateMission(mission models.Mission) (models.Mission, error) {
	result := database.DB.Create(&mission)
	if result.Error != nil {
		return mission, result.Error
	}

	// Reload related alchemist
	if err := database.DB.Preload("Alchemist").First(&mission, mission.ID).Error; err != nil {
		return mission, err
	}

	return mission, nil
}

// GetAllMissions retrieves all missions from the database
func GetAllMissions() ([]models.Mission, error) {
	var missions []models.Mission
	result := database.DB.Preload("Alchemist").Find(&missions)
	return missions, result.Error
}

// GetMissionByID retrieves a single mission by ID
func GetMissionByID(id uint) (models.Mission, error) {
	var mission models.Mission
	result := database.DB.Preload("Alchemist").First(&mission, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return mission, errors.New("mission not found")
	}
	return mission, result.Error
}

// UpdateMission updates an existing mission by ID
func UpdateMission(id uint, updated models.Mission) (models.Mission, error) {
	var mission models.Mission
	if err := database.DB.First(&mission, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return mission, errors.New("mission not found")
		}
		return mission, err
	}

	mission.Title = updated.Title
	mission.Description = updated.Description
	mission.Status = updated.Status
	mission.AlchemistID = updated.AlchemistID

	if err := database.DB.Save(&mission).Error; err != nil {
		return mission, err
	}

	// Reload mission with its related Alchemist
	if err := database.DB.Preload("Alchemist").First(&mission, mission.ID).Error; err != nil {
		return mission, err
	}

	return mission, nil
}

// DeleteMission removes a mission by ID
func DeleteMission(id uint) error {
	var mission models.Mission
	if err := database.DB.First(&mission, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("mission not found")
		}
		return err
	}
	return database.DB.Delete(&mission).Error
}
