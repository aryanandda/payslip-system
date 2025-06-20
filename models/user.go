package models

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	IsAdmin      bool   `gorm:"default:false"`
	Salary       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
