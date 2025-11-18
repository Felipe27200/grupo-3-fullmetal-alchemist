package models

import "time"

type Transmutation struct {
    ID          uint       `gorm:"primaryKey" json:"id"`
    Input       string     `json:"input"`
    Output      string     `json:"output"`
    Approved    bool       `json:"approved"`
    Status      string     `json:"status"`
    Result      string     `json:"result"`
    Error       string     `json:"error,omitempty"`
    ExecutedAt  *time.Time `json:"executed_at"`

    AlchemistID uint       `json:"alchemist_id"`
    Alchemist   Alchemist  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"alchemist,omitempty"`
}
