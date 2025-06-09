package mocks

import (
	"payslip-system/models"
)

type MockPayrollController struct {
	CreatePayrollFunc func(periodID, adminID uint) (uint, error)
	GetPayrollFunc    func(payrollID uint) (*models.Payroll, error)
}

func (m *MockPayrollController) CreatePayroll(pid, aid uint) (uint, error) {
	return m.CreatePayrollFunc(pid, aid)
}

func (m *MockPayrollController) GetPayroll(payrollID uint) (*models.Payroll, error) {
	return m.GetPayrollFunc(payrollID)
}
