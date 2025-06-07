package controllers

import (
	"time"

	"gorm.io/gorm"
	"payslip-system/models"
)

type AttendanceController struct {
	DB *gorm.DB
}

func NewAttendanceController(db *gorm.DB) *AttendanceController {
	return &AttendanceController{DB: db}
}

// Check if attendance exists for user on date
func (ctrl *AttendanceController) GetAttendanceByUserAndDate(userID uint, date time.Time) (*models.Attendance, error) {
	var attendance models.Attendance
	err := ctrl.DB.Where("user_id = ? AND date = ?", userID, date.Format("2006-01-02")).First(&attendance).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &attendance, nil
}

// Create new attendance record
func (ctrl *AttendanceController) CreateAttendance(attendance *models.Attendance) error {
	return ctrl.DB.Create(attendance).Error
}
