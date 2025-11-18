package services

import (
	"alchemy-system/database"
	"alchemy-system/models"
	"fmt"
)

func GetAllMaterials() ([]models.Material, error) {
	var materials []models.Material
	err := database.DB.Find(&materials).Error
	return materials, err
}

func GetMaterialByID(id string) (models.Material, error) {
	var material models.Material
	err := database.DB.First(&material, id).Error
	return material, err
}

func CreateMaterial(mat models.Material) (models.Material, error) {
	if err := database.DB.Create(&mat).Error; err != nil {
		return mat, err
	}

	// Auditoría: material creado
	CreateAudit(models.Audit{
		Action:   "CREATE",
		Entity:   "material",
		EntityID: mat.ID,
		Message:  fmt.Sprintf("Material %s created with qty=%d", mat.Name, mat.Quantity),
	})

	return mat, nil
}

func UpdateMaterial(id string, data models.Material) (models.Material, error) {
	var material models.Material

	if err := database.DB.First(&material, id).Error; err != nil {
		return material, err
	}

	material.Name = data.Name
	material.Quantity = data.Quantity
	material.Danger = data.Danger

	if err := database.DB.Save(&material).Error; err != nil {
		return material, err
	}

	// Auditoría: material actualizado
	CreateAudit(models.Audit{
		Action:   "UPDATE",
		Entity:   "material",
		EntityID: material.ID,
		Message:  fmt.Sprintf("Material %s updated (qty=%d)", material.Name, material.Quantity),
	})

	return material, nil
}

func DeleteMaterial(id string) error {
	var material models.Material

	if err := database.DB.First(&material, id).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&material).Error; err != nil {
		return err
	}

	// Auditoría: material eliminado
	CreateAudit(models.Audit{
		Action:   "DELETE",
		Entity:   "material",
		EntityID: material.ID,
		Message:  fmt.Sprintf("Material %s deleted", material.Name),
	})

	return nil
}
