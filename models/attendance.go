package models

import "time"

type Attendance struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Date      time.Time `gorm:"type:date"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
