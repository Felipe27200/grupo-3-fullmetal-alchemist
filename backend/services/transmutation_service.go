	package services

	import (
		"alchemy-system/database"
    	"alchemy-system/models"
		"errors"

		"gorm.io/gorm"
	)

	// CreateTransmutation inserts a new transmutation into the database and preloads its Alchemist.
	func CreateTransmutation(transmutation models.Transmutation) (models.Transmutation, error) {
		result := database.DB.Create(&transmutation)
		if result.Error != nil {
			return transmutation, result.Error
		}

		// Reload related alchemist
		if err := database.DB.Preload("Alchemist").First(&transmutation, transmutation.ID).Error; err != nil {
			return transmutation, err
		}

		return transmutation, nil
	}

	// GetAllTransmutations retrieves all transmutations from the database (with their Alchemists).
	func GetAllTransmutations() ([]models.Transmutation, error) {
		var transmutations []models.Transmutation
		result := database.DB.Preload("Alchemist").Find(&transmutations)
		return transmutations, result.Error
	}

	// GetTransmutationByID retrieves a transmutation by its ID (with its Alchemist).
	func GetTransmutationByID(id uint) (models.Transmutation, error) {
		var transmutation models.Transmutation
		result := database.DB.Preload("Alchemist").First(&transmutation, id)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return transmutation, errors.New("transmutation not found")
		}
		return transmutation, result.Error
	}

	// UpdateTransmutation updates an existing transmutation by ID and returns the updated one (with Alchemist).
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

		// Reload updated transmutation with its Alchemist
		if err := database.DB.Preload("Alchemist").First(&transmutation, transmutation.ID).Error; err != nil {
			return transmutation, err
		}

		return transmutation, nil
	}

	// DeleteTransmutation removes a transmutation by ID.
	func DeleteTransmutation(id uint) error {
		var transmutation models.Transmutation
		if err := database.DB.First(&transmutation, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("transmutation not found")
			}
			return err
		}

		return database.DB.Delete(&transmutation).Error
	}
