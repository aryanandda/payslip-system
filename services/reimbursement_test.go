package services

import (
	"errors"
	"payslip-system/constants"
	"payslip-system/mocks"
	"testing"
)

func TestSubmitReimbursement_Success(t *testing.T) {
	mock := &mocks.MockReimbursementController{
		CreateReimbursementFunc: func(userID uint, amount float64, description string) error {
			if userID != 1 || amount != 100000 || description != "Lunch" {
				t.Errorf("unexpected input to CreateReimbursement")
			}
			return nil
		},
	}
	service := NewReimbursementService(mock)

	err := service.SubmitReimbursement(1, 100000, "Lunch")
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestSubmitReimbursement_InvalidAmount(t *testing.T) {
	service := NewReimbursementService(nil)

	err := service.SubmitReimbursement(1, 0, "Lunch")
	if err == nil || err.Error() != constants.ErrInvalidAmount {
		t.Errorf("expected error %v, got %v", constants.ErrInvalidAmount, err)
	}
}

func TestSubmitReimbursement_MissingDescription(t *testing.T) {
	service := NewReimbursementService(nil)

	err := service.SubmitReimbursement(1, 100000, "")
	if err == nil || err.Error() != constants.ErrMissingDescription {
		t.Errorf("expected error %v, got %v", constants.ErrMissingDescription, err)
	}
}

func TestSubmitReimbursement_CreateFails(t *testing.T) {
	mock := &mocks.MockReimbursementController{
		CreateReimbursementFunc: func(userID uint, amount float64, description string) error {
			return errors.New("DB failure")
		},
	}
	service := NewReimbursementService(mock)

	err := service.SubmitReimbursement(1, 100000, "Transport")
	if err == nil || err.Error() != "DB failure" {
		t.Errorf("expected DB failure error, got %v", err)
	}
}
