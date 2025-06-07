package models

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID        uint      `gorm:"primaryKey"`
	TableName string    `gorm:"not null"`
	Action    string    `gorm:"not null"`
	RecordID  uint      `gorm:"not null"`
	UserID    *uint     `gorm:"index"`
	IPAddress string    `gorm:"size:45"`
	RequestID uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
