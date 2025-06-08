package controllers

import (
	"errors"

	"gorm.io/gorm"

	"payslip-system/models"
)

type PayrollController struct {
	DB *gorm.DB
}

func NewPayrollController(db *gorm.DB) *PayrollController {
	return &PayrollController{DB: db}
}

func (c *PayrollController) CreatePayroll(attendancePeriodID, adminID uint) (uint, error) {
	payroll := models.Payroll{
		AttendancePeriodID: attendancePeriodID,
		CreatedBy:          &adminID,
	}

	if err := c.DB.Create(&payroll).Error; err != nil {
		return 0, err
	}

	return payroll.ID, nil
}

func (ctrl *PayrollController) GetPayroll(payrollID uint) (*models.Payroll, error) {
	var payroll models.Payroll
	if err := ctrl.DB.Where("id = ?", payrollID).First(&payroll).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &payroll, nil
}
