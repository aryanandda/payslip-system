package services

import (
	"errors"
	"payslip-system/constants"
	"payslip-system/controllers"
)

type ReimbursementService struct {
	ctrl *controllers.ReimbursementController
}

func NewReimbursementService(ctrl *controllers.ReimbursementController) *ReimbursementService {
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
