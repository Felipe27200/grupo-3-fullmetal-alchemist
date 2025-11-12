package models

import "time"

type Audit struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Action    string    `json:"action"`
	Entity    string    `json:"entity"`
	EntityID  uint      `json:"entity_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
