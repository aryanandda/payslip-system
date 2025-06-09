package services

import (
	"errors"
	"payslip-system/constants"
	"payslip-system/interfaces"
)

type ReimbursementService struct {
	ctrl interfaces.ReimbursementControllerInterface
}

func NewReimbursementService(ctrl interfaces.ReimbursementControllerInterface) *ReimbursementService {
	return &ReimbursementService{ctrl: ctrl}
}

func (s *ReimbursementService) SubmitReimbursement(userID uint, amount float64, description string) error {
	if amount <= 0 {
		return errors.New(constants.ErrInvalidAmount)
	}
	if description == "" {
		return errors.New(constants.ErrMissingDescription)
	}
	return s.ctrl.CreateReimbursement(userID, amount, description)
}
