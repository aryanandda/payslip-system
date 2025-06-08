package models

import "time"

type Payroll struct {
	ID                 uint `gorm:"primaryKey"`
	AttendancePeriodID uint `gorm:"foreignKey"`
	CreatedBy          *uint
	CreatedAt          time.Time
}
