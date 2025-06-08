package models

import "time"

type Attendance struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Date      time.Time `gorm:"type:date"`
	CreatedAt time.Time
}

type AttendancePeriod struct {
	ID        uint      `gorm:"primaryKey"`
	StartDate time.Time `gorm:"type:date;not null"`
	EndDate   time.Time `gorm:"type:date;not null"`
	IsClosed  bool      `gorm:"default:false"`
	CreatedBy *uint
	CreatedAt time.Time
	UpdatedBy *uint
	UpdatedAt time.Time
}
