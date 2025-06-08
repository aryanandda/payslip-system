package models

import "time"

type Overtime struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Date      time.Time `gorm:"not null;index"`
	Hours     float64   `gorm:"not null"`
	CreatedAt time.Time
}
