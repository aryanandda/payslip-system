package mocks

import (
	"payslip-system/models"
	"time"
)

type MockReimbursementController struct {
	CreateReimbursementFunc func(userID uint, amount float64, description string) error
	GetReimbursementsFunc   func(userID uint, start, end time.Time) ([]models.Reimbursement, error)
}

func (m *MockReimbursementController) CreateReimbursement(userID uint, amount float64, description string) error {
	return m.CreateReimbursementFunc(userID, amount, description)
}

func (m *MockReimbursementController) GetReimbursements(userID uint, start, end time.Time) ([]models.Reimbursement, error) {
	return m.GetReimbursementsFunc(userID, start, end)
}
