package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Clave secreta para firmar JWT
// Cámbiala por una variable de entorno en producción
var JwtSecret = []byte("MI_SUPER_SECRETO_QUE_DEBES_CAMBIAR")

// GenerateJWT crea un token JWT válido por 24 horas
func GenerateJWT(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	// Crear token usando HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar token con la clave secreta
	return token.SignedString(JwtSecret)
}
