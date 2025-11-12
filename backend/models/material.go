package models

type Material struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	Danger    string `json:"danger"`
}
