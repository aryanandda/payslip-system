package interfaces

import (
	"payslip-system/models"
)

type UserControllerInterface interface {
	GetAllEmployee() ([]models.User, error)
	FindByUsername(username string) (*models.User, error)
}
