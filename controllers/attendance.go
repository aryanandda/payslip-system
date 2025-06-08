package controllers

import (
	"time"

	"payslip-system/models"

	"gorm.io/gorm"
)

type AttendanceController struct {
	DB *gorm.DB
}

func NewAttendanceController(db *gorm.DB) *AttendanceController {
	return &AttendanceController{DB: db}
}

func (ctrl *AttendanceController) CheckOverlap(start, end time.Time) (bool, error) {
	var count int64
	err := ctrl.DB.Model(&models.AttendancePeriod{}).
		Where("start_date <= ? AND end_date >= ?", end, start).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ctrl *AttendanceController) CreateAttendancePeriod(period *models.AttendancePeriod) error {
	return ctrl.DB.Create(period).Error
}

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

func (ctrl *AttendanceController) GetAttendanceByUserAndDateBetween(userID uint, startDate, endDate time.Time) ([]models.Attendance, error) {
	attendances := make([]models.Attendance, 0)
	err := ctrl.DB.Where("user_id = ? AND date BETWEEN ? AND ?", userID, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).Find(&attendances).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return attendances, nil
}

func (ctrl *AttendanceController) GetAttendancePeriodByID(id uint) (*models.AttendancePeriod, error) {
	var attendancePeriod models.AttendancePeriod
	err := ctrl.DB.Where("id = ?", id).First(&attendancePeriod).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &attendancePeriod, nil
}

func (ctrl *AttendanceController) CreateAttendance(attendance *models.Attendance) error {
	return ctrl.DB.Create(attendance).Error
}

func (c *AttendanceController) CountAttendance(userID uint, start, end time.Time) (int64, error) {
	var count int64
	err := c.DB.Model(&models.Attendance{}).
		Where("user_id = ? AND date BETWEEN ? AND ?", userID, start, end).
		Count(&count).Error
	return count, err
}

func (c *AttendanceController) ClosePeriod(id uint, userID uint) error {
	var period models.AttendancePeriod
	if err := c.DB.First(&period, id).Error; err != nil {
		return err
	}

	period.IsClosed = true
	period.UpdatedBy = &userID
	return c.DB.Save(&period).Error
}
