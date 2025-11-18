package services


import (
	"alchemy-system/database"
	"alchemy-system/models"
	"alchemy-system/queue"
	"errors"

	"gorm.io/gorm"
)

func CreateTransmutation(transmutation models.Transmutation) (models.Transmutation, error) {
	if err := database.DB.Create(&transmutation).Error; err != nil {
		return transmutation, err
	}

	if err := database.DB.Preload("Alchemist").First(&transmutation, transmutation.ID).Error; err != nil {
		return transmutation, err
	}

	// Publicar tarea en la cola
	queue.PublishTransmutation(map[string]any{
		"id":           transmutation.ID,
		"alchemist_id": transmutation.AlchemistID,
		"input":        transmutation.Input,
	})

	// Auditoría
	CreateAudit(models.Audit{
		Action:   "CREATE",
		Entity:   "transmutation",
		EntityID: transmutation.ID,
		Message:  "New transmutation created and queued",
	})

	return transmutation, nil
}

func GetAllTransmutations() ([]models.Transmutation, error) {
	var transmutations []models.Transmutation
	err := database.DB.Preload("Alchemist").Find(&transmutations).Error
	return transmutations, err
}

func GetTransmutationByID(id uint) (models.Transmutation, error) {
	var transmutation models.Transmutation

	err := database.DB.Preload("Alchemist").First(&transmutation, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return transmutation, errors.New("transmutation not found")
	}

	return transmutation, err
}

func UpdateTransmutation(id uint, updated models.Transmutation) (models.Transmutation, error) {
	var transmutation models.Transmutation

	if err := database.DB.First(&transmutation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return transmutation, errors.New("transmutation not found")
		}
		return transmutation, err
	}

	transmutation.Input = updated.Input
	transmutation.Output = updated.Output
	transmutation.Approved = updated.Approved
	transmutation.ExecutedAt = updated.ExecutedAt
	transmutation.AlchemistID = updated.AlchemistID

	if err := database.DB.Save(&transmutation).Error; err != nil {
		return transmutation, err
	}

	if err := database.DB.Preload("Alchemist").First(&transmutation, transmutation.ID).Error; err != nil {
		return transmutation, err
	}

	// Auditoría
	CreateAudit(models.Audit{
		Action:   "UPDATE",
		Entity:   "transmutation",
		EntityID: transmutation.ID,
		Message:  "Transmutation updated",
	})

	return transmutation, nil
}

func DeleteTransmutation(id uint) error {
	var transmutation models.Transmutation

	if err := database.DB.First(&transmutation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("transmutation not found")
		}
		return err
	}

	if err := database.DB.Delete(&transmutation).Error; err != nil {
		return err
	}

	// Auditoría
	CreateAudit(models.Audit{
		Action:   "DELETE",
		Entity:   "transmutation",
		EntityID: transmutation.ID,
		Message:  "Transmutation deleted",
	})

	return nil
}
