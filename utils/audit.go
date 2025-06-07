package utils

import (
	"time"

	"payslip-system/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func LogAudit(db *gorm.DB, table, action string, recordID uint, userID uint, ip string, reqID string) error {
	parsedReqID, _ := uuid.Parse(reqID)
	log := models.AuditLog{
		TableName: table,
		Action:    action,
		RecordID:  recordID,
		UserID:    &userID,
		IPAddress: ip,
		RequestID: parsedReqID,
		CreatedAt: time.Now(),
	}
	return db.Create(&log).Error
}
