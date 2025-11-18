package services

import (
	"alchemy-system/database"
	"alchemy-system/models"
)

func CreateAudit(a models.Audit) (models.Audit, error) {
	err := database.DB.Create(&a).Error
	return a, err
}

func GetAllAudits() ([]models.Audit, error) {
	var audits []models.Audit
	err := database.DB.Find(&audits).Error
	return audits, err
}

func GetAuditByID(id uint) (models.Audit, error) {
	var audit models.Audit
	err := database.DB.First(&audit, id).Error
	return audit, err
}

func DeleteAudit(id uint) error {
	return database.DB.Delete(&models.Audit{}, id).Error
}
