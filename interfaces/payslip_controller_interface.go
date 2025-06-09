package interfaces

import (
	"payslip-system/models"
)

type PayslipControllerInterface interface {
	CreatePayslip(payslip models.Payslip) error
	GetPayslipByPayrollID(userID, payrollID uint) (*models.Payslip, error)
	GetPayslips(payrollID uint) ([]models.Payslip, error)
}
