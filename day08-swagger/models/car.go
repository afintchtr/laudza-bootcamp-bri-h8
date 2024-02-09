package models

import "time"

// Car represents the model for an car
type Car struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(191)"`
	Price     uint   `gorm:"not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}