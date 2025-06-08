package controllers

import (
	"payslip-system/models"
	"time"

	"gorm.io/gorm"
)

type OvertimeController struct {
	DB *gorm.DB
}

func NewOvertimeController(db *gorm.DB) *OvertimeController {
	return &OvertimeController{DB: db}
}

func (c *OvertimeController) CreateOvertime(userID uint, date time.Time, hours float64) error {
	overtime := models.Overtime{
		UserID: userID,
		Date:   date,
		Hours:  hours,
	}
	return c.DB.Create(&overtime).Error
}

func (c *OvertimeController) GetOvertimesByDate(userID uint, date time.Time) ([]float64, error) {
	var total []float64
	err := c.DB.Model(&models.Overtime{}).
		Where("user_id = ? AND DATE(date) = ?", userID, date.Format("2006-01-02")).
		Select("hours").Scan(&total).Error
	return total, err
}

func (c *OvertimeController) GetOvertimeTotal(userID uint, start, end time.Time) (float64, error) {
	var total float64
	err := c.DB.Model(&models.Overtime{}).
		Where("user_id = ? AND date BETWEEN ? AND ?", userID, start, end).
		Select("SUM(hours)").Scan(&total).Error
	return total, err
}

func (c *OvertimeController) GetOvertimeGroupedByDate(userID uint, start, end time.Time) ([]struct {
	Date  time.Time
	Hours float64
}, error) {
	var results []struct {
		Date  time.Time
		Hours float64
	}

	err := c.DB.Model(&models.Overtime{}).
		Select("DATE(date), SUM(hours) AS hours").
		Where("user_id = ? AND date BETWEEN ? AND ?", userID, start, end).
		Group("DATE(date)").
		Order("date").
		Scan(&results).Error

	return results, err
}
