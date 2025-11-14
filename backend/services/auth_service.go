package services

import (
	"alchemy-system/database"
	"alchemy-system/models"
	"alchemy-system/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user models.Alchemist) (models.Alchemist, error) {

	// Encriptar password
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hashed)

	// Guardar en la DB
	if err := database.DB.Create(&user).Error; err != nil {
		return user, err
	}

	// Nunca devolver password
	user.Password = ""

	return user, nil
}

func Login(email, password string) (string, error) {
	var user models.Alchemist

	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	// Verificar password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("incorrect password")
	}

	// Crear token
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
