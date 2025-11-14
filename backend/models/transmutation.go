package models

import "time"

type Transmutation struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	Input       string      `json:"input"`
	Output      string      `json:"output"`
	Approved    bool        `json:"approved"`
	ExecutedAt  *time.Time  `json:"executed_at"`  // ← CORRECCIÓN: debe ser puntero

	AlchemistID uint        `json:"alchemist_id"`
	Alchemist   Alchemist   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"alchemist,omitempty"`
}
