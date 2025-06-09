package mocks

import (
	"payslip-system/models"
)

type MockPayslipController struct {
	CreatePayslipFunc         func(payslip models.Payslip) error
	GetPayslipByPayrollIDFunc func(userID, payrollID uint) (*models.Payslip, error)
	GetPayslipsFunc           func(payrollID uint) ([]models.Payslip, error)
}

func (m *MockPayslipController) CreatePayslip(payslip models.Payslip) error {
	return m.CreatePayslipFunc(payslip)
}

func (m *MockPayslipController) GetPayslipByPayrollID(userID, payrollID uint) (*models.Payslip, error) {
	return m.GetPayslipByPayrollIDFunc(userID, payrollID)
}

func (m *MockPayslipController) GetPayslips(payrollID uint) ([]models.Payslip, error) {
	return m.GetPayslipsFunc(payrollID)
}
