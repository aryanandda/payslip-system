package controllers

import (
	"time"

	"payslip-system/models"

	"gorm.io/gorm"
)

type PayrollController struct {
	DB *gorm.DB
}

func NewPayrollController(db *gorm.DB) *PayrollController {
	return &PayrollController{DB: db}
}

func (ctrl *PayrollController) CheckOverlap(start, end time.Time) (bool, error) {
	var count int64
	err := ctrl.DB.Model(&models.PayrollPeriod{}).
		Where("start_date <= ? AND end_date >= ?", end, start).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ctrl *PayrollController) CreatePayrollPeriod(period *models.PayrollPeriod) error {
	return ctrl.DB.Create(period).Error
}
