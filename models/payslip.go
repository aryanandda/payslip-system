package models

import "time"

type Payslip struct {
	ID                 uint `gorm:"primaryKey"`
	UserID             uint `gorm:"foreignKey"`
	PayrollID          uint `gorm:"foreignKey"`
	OvertimeHours      float64
	OvertimePay        float64
	ReimbursementTotal float64
	TotalTakeHome      float64
	RateSalaryPerDay   float64
	RateSalaryPerHour  float64
	PresentDays        int
	CreatedBy          *uint
	CreatedAt          time.Time
	User               User `gorm:"foreignKey:UserID"`
}
