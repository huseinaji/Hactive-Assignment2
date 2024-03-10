package models

import (
	"time"
)

type Order struct {
	ID           uint   `gorm:"primaryKey"`
	CustomerName string `gorm:"not null;type:varchar(100)"`
	OrderedAt    string `gorm:"not null;type:varchar(100)"`
	Items        []Item
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
