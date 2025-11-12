package models

import "time"

type Mission struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`

	// Relation to Alchemist
	AlchemistID uint      `json:"alchemist_id"`
	Alchemist   Alchemist `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"alchemist,omitempty"`
}
