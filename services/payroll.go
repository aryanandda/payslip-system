package services

import (
	"errors"
	"time"

	"payslip-system/controllers"
	"payslip-system/models"
)

type PayrollService struct {
	ctrl *controllers.PayrollController
}

func NewPayrollService(ctrl *controllers.PayrollController) *PayrollService {
	return &PayrollService{ctrl: ctrl}
}

func (s *PayrollService) ValidateAndCreatePeriod(startStr, endStr string, userID uint) (*models.PayrollPeriod, error) {
	startDate, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		return nil, errors.New("invalid start_date format")
	}
	endDate, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		return nil, errors.New("invalid end_date format")
	}
	if !startDate.Before(endDate) {
		return nil, errors.New("start_date must be before end_date")
	}

	// Check for overlap via controller
	overlap, err := s.ctrl.CheckOverlap(startDate, endDate)
	if err != nil {
		return nil, err
	}
	if overlap {
		return nil, errors.New("overlapping payroll period exists")
	}

	// Prepare model
	return &models.PayrollPeriod{
		StartDate: startDate,
		EndDate:   endDate,
		CreatedBy: &userID,
	}, nil
}
