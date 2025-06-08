package models

import "time"

type Reimbursement struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"not null"`
	Amount      float64 `gorm:"not null"`
	Description string  `gorm:"type:text;not null"`
	CreatedAt   time.Time
}
