package interfaces

import (
	"payslip-system/models"
)

type PayrollControllerInterface interface {
	CreatePayroll(attendancePeriodID, adminID uint) (uint, error)
	GetPayroll(payrollID uint) (*models.Payroll, error)
}
