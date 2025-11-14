package models

import "time"

type Alchemist struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Name       string     `json:"name"`
	Email      string     `gorm:"unique;not null" json:"email"`
	Password   string     `json:"password,omitempty"` // Se recibe en el request pero NO se envía en las respuestas
	Rank       string     `json:"rank"`
	Specialty  string     `json:"specialty"`
	Role       string     `json:"role"` // Ej: "alchemist", "supervisor"

	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`

	Missions   []Mission  `gorm:"foreignKey:AlchemistID" json:"missions,omitempty"`
}
