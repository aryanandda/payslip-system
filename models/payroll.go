package models

import "time"

type PayrollPeriod struct {
	ID        uint      `gorm:"primaryKey"`
	StartDate time.Time `gorm:"type:date;not null"`
	EndDate   time.Time `gorm:"type:date;not null"`
	IsClosed  bool      `gorm:"default:false"`
	CreatedBy *uint
	UpdatedBy *uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
