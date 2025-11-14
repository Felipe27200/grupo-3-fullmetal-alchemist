package services

import (
	"alchemy-system/backend/database"
	"alchemy-system/backend/models"
)

func GetAllMaterials() ([]models.Material, error) {
	var materials []models.Material
	result := database.DB.Find(&materials)
	return materials, result.Error
}

func GetMaterialByID(id string) (models.Material, error) {
	var material models.Material
	result := database.DB.First(&material, id)
	return material, result.Error
}

func CreateMaterial(mat models.Material) (models.Material, error) {
	result := database.DB.Create(&mat)
	return mat, result.Error
}

func UpdateMaterial(id string, data models.Material) (models.Material, error) {
	var material models.Material

	if err := database.DB.First(&material, id).Error; err != nil {
		return material, err
	}

	material.Name = data.Name
	material.Quantity = data.Quantity
	material.Danger = data.Danger

	database.DB.Save(&material)
	return material, nil
}

func DeleteMaterial(id string) error {
	return database.DB.Delete(&models.Material{}, id).Error
}
