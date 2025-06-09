package interfaces

import (
	"payslip-system/models"
	"time"
)

type ReimbursementControllerInterface interface {
	CreateReimbursement(userID uint, amount float64, description string) error
	GetReimbursements(userID uint, start, end time.Time) ([]models.Reimbursement, error)
}
