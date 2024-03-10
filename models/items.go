package models

import "time"

type Item struct {
	ID          uint   `gorm:"primaryKey"`
	ItemCode    string `gorm:"not null;type:varchar(100)"`
	Description string `gorm:"not null;type:varchar(100)"`
	Quantity    int
	OrderID     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
