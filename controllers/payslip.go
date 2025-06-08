package controllers

import (
	"gorm.io/gorm"

	"payslip-system/models"
)

type PayslipController struct {
	DB *gorm.DB
}

func NewPayslipController(db *gorm.DB) *PayslipController {
	return &PayslipController{DB: db}
}

func (c *PayslipController) CreatePayslip(payslip models.Payslip) error {
	return c.DB.Create(&payslip).Error
}

func (c *PayslipController) GetPayslipByPayrollID(userID, payrollID uint) (*models.Payslip, error) {
	var payslip models.Payslip
	err := c.DB.Where("user_id = ? AND payroll_id = ?", userID, payrollID).
		First(&payslip).Error
	return &payslip, err
}

func (c *PayslipController) GetPayslips(payrollID uint) ([]models.Payslip, error) {
	payslips := make([]models.Payslip, 0)
	err := c.DB.Preload("User").
		Where("payroll_id = ?", payrollID).
		Find(&payslips).Error

	return payslips, err
}
