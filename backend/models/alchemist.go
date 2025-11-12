package models

import "time"

type Alchemist struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	Rank       string    `json:"rank"`
	Specialty  string    `json:"specialty"`
	CreatedAt  time.Time `json:"created_at"`
	Missions   []Mission `gorm:"foreignKey:AlchemistID" json:"missions,omitempty"`
}
