package services

import (
	"alchemy-system/database"
	"alchemy-system/models"
	"errors"

	"gorm.io/gorm"
)

func CreateMission(mission models.Mission) (models.Mission, error) {
	if err := database.DB.Create(&mission).Error; err != nil {
		return mission, err
	}

	if err := database.DB.Preload("Alchemist").First(&mission, mission.ID).Error; err != nil {
		return mission, err
	}

	// Auditoría: misión creada
	CreateAudit(models.Audit{
		Action:   "CREATE",
		Entity:   "mission",
		EntityID: mission.ID,
		Message:  "Mission created",
	})

	return mission, nil
}

func GetAllMissions() ([]models.Mission, error) {
	var missions []models.Mission
	err := database.DB.Preload("Alchemist").Find(&missions).Error
	return missions, err
}

func GetMissionByID(id uint) (models.Mission, error) {
	var mission models.Mission

	err := database.DB.Preload("Alchemist").First(&mission, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return mission, errors.New("mission not found")
	}

	return mission, err
}

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

	if err := database.DB.Preload("Alchemist").First(&mission, mission.ID).Error; err != nil {
		return mission, err
	}

	// Auditoría: misión actualizada
	CreateAudit(models.Audit{
		Action:   "UPDATE",
		Entity:   "mission",
		EntityID: mission.ID,
		Message:  "Mission updated",
	})

	return mission, nil
}

func DeleteMission(id uint) error {
	var mission models.Mission

	err := database.DB.First(&mission, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("mission not found")
	}
	if err != nil {
		return err
	}

	if err := database.DB.Delete(&mission).Error; err != nil {
		return err
	}

	// Auditoría: misión eliminada
	CreateAudit(models.Audit{
		Action:   "DELETE",
		Entity:   "mission",
		EntityID: id,
		Message:  "Mission deleted",
	})

	return nil
}
