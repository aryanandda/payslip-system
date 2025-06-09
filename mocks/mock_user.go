package mocks

import (
	"payslip-system/models"
)

type MockUserController struct {
	GetAllEmployeeFunc func() ([]models.User, error)
	FindByUsernameFunc func(username string) (*models.User, error)
}

func (m *MockUserController) GetAllEmployee() ([]models.User, error) {
	return m.GetAllEmployeeFunc()
}

func (m *MockUserController) FindByUsername(username string) (*models.User, error) {
	return m.FindByUsernameFunc(username)
}
