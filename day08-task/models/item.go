package models

import "time"

// Item represents the model for an item in order
type Item struct {
	ItemID      uint   `gorm:"primaryKey"`
	ItemCode    string `gorm:"not null;type:varchar(255)"`
	Description string `gorm:"not null;type:varchar(255)"`
	Quantity    uint   `gorm:"not null;"`
	OrderID     uint	 `gorm:"not null;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}