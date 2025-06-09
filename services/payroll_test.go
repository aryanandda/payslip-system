package services

import (
	"errors"
	"payslip-system/dto"
	"payslip-system/mocks"
	"payslip-system/models"
	"testing"
	"time"
)

func TestRunPayroll_Success(t *testing.T) {
	attendancePeriod := &models.AttendancePeriod{
		ID:        1,
		IsClosed:  false,
		StartDate: time.Date(2025, 5, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2025, 5, 31, 0, 0, 0, 0, time.UTC),
	}

	mockService := NewPayrollService(
		&mocks.MockPayrollController{
			CreatePayrollFunc: func(pid, aid uint) (uint, error) { return 101, nil },
		},
		&mocks.MockAttendanceController{
			GetAttendancePeriodByIDFunc: func(id uint) (*models.AttendancePeriod, error) {
				return attendancePeriod, nil
			},
			ClosePeriodFunc: func(id uint, aid uint) error { return nil },
			GetAttendanceByUserAndDateBetweenFunc: func(uid uint, start, end time.Time) ([]models.Attendance, error) {
				return []models.Attendance{{UserID: uid}}, nil
			},
		},
		&mocks.MockUserController{
			GetAllEmployeeFunc: func() ([]models.User, error) {
				return []models.User{{ID: 1, Salary: 2200000}}, nil
			},
		},
		&mocks.MockOvertimeController{
			GetOvertimeTotalFunc: func(uid uint, start, end time.Time) (float64, error) {
				return 2.0, nil
			},
		},
		&mocks.MockReimbursementController{
			GetReimbursementsFunc: func(uid uint, start, end time.Time) ([]models.Reimbursement, error) {
				return []models.Reimbursement{{Amount: 100000}}, nil
			},
		},
		&mocks.MockPayslipController{
			CreatePayslipFunc: func(p models.Payslip) error {
				if p.UserID != 1 || p.TotalTakeHome <= 0 {
					t.Errorf("Invalid payslip data: %+v", p)
				}
				return nil
			},
		},
	)

	req := dto.RunPayrollRequest{AttendancePeriodID: 1}
	err := mockService.RunPayroll(req, 99)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestRunPayroll_PeriodAlreadyClosed(t *testing.T) {
	mockService := NewPayrollService(
		&mocks.MockPayrollController{},
		&mocks.MockAttendanceController{
			GetAttendancePeriodByIDFunc: func(id uint) (*models.AttendancePeriod, error) {
				return &models.AttendancePeriod{
					ID:       1,
					IsClosed: true,
				}, nil
			},
		},
		&mocks.MockUserController{},
		&mocks.MockOvertimeController{},
		&mocks.MockReimbursementController{},
		&mocks.MockPayslipController{},
	)

	req := dto.RunPayrollRequest{AttendancePeriodID: 1}
	err := mockService.RunPayroll(req, 99)

	if err == nil || err.Error() != "payroll already processed for this period" {
		t.Fatalf("expected closed period error, got: %v", err)
	}
}

func TestRunPayroll_ClosePeriodError(t *testing.T) {
	attendancePeriod := &models.AttendancePeriod{
		ID:        1,
		IsClosed:  false,
		StartDate: time.Date(2025, 5, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2025, 5, 31, 0, 0, 0, 0, time.UTC),
	}

	mockService := NewPayrollService(
		&mocks.MockPayrollController{},
		&mocks.MockAttendanceController{
			GetAttendancePeriodByIDFunc: func(id uint) (*models.AttendancePeriod, error) {
				return attendancePeriod, nil
			},
			ClosePeriodFunc: func(id uint, aid uint) error {
				return errors.New("db connection failed")
			},
		},
		&mocks.MockUserController{},
		&mocks.MockOvertimeController{},
		&mocks.MockReimbursementController{},
		&mocks.MockPayslipController{},
	)

	req := dto.RunPayrollRequest{AttendancePeriodID: 1}
	err := mockService.RunPayroll(req, 99)

	if err == nil || err.Error() != "error when close period" {
		t.Fatalf("expected close period error, got: %v", err)
	}
}

func TestRunPayroll_GetAttendancePeriodByIDError(t *testing.T) {
	mockService := NewPayrollService(
		&mocks.MockPayrollController{},
		&mocks.MockAttendanceController{
			GetAttendancePeriodByIDFunc: func(id uint) (*models.AttendancePeriod, error) {
				return nil, errors.New("failed to fetch period")
			},
		},
		&mocks.MockUserController{},
		&mocks.MockOvertimeController{},
		&mocks.MockReimbursementController{},
		&mocks.MockPayslipController{},
	)

	req := dto.RunPayrollRequest{AttendancePeriodID: 1}
	err := mockService.RunPayroll(req, 42)

	if err == nil || err.Error() != "failed to fetch period" {
		t.Fatalf("expected fetch period error, got: %v", err)
	}
}

func TestRunPayroll_CreatePayrollError(t *testing.T) {
	attendancePeriod := &models.AttendancePeriod{
		ID:        1,
		IsClosed:  false,
		StartDate: time.Date(2025, 5, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2025, 5, 31, 0, 0, 0, 0, time.UTC),
	}

	mockService := NewPayrollService(
		&mocks.MockPayrollController{
			CreatePayrollFunc: func(pid, aid uint) (uint, error) {
				return 0, errors.New("error when create payroll")
			},
		},
		&mocks.MockAttendanceController{
			GetAttendancePeriodByIDFunc: func(id uint) (*models.AttendancePeriod, error) {
				return attendancePeriod, nil
			},
			ClosePeriodFunc: func(id uint, aid uint) error { return nil },
		},
		&mocks.MockUserController{},
		&mocks.MockOvertimeController{},
		&mocks.MockReimbursementController{},
		&mocks.MockPayslipController{},
	)

	req := dto.RunPayrollRequest{AttendancePeriodID: 1}
	err := mockService.RunPayroll(req, 99)
	if err == nil || err.Error() != "error when create payroll" {
		t.Fatalf("expected create payroll error, got: %v", err)
	}
}

func TestRunPayroll_GetAllEmployeeError(t *testing.T) {
	attendancePeriod := &models.AttendancePeriod{
		ID:        1,
		IsClosed:  false,
		StartDate: time.Date(2025, 5, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2025, 5, 31, 0, 0, 0, 0, time.UTC),
	}

	mockService := NewPayrollService(
		&mocks.MockPayrollController{
			CreatePayrollFunc: func(pid, aid uint) (uint, error) {
				return 1, nil
			},
		},
		&mocks.MockAttendanceController{
			GetAttendancePeriodByIDFunc: func(id uint) (*models.AttendancePeriod, error) {
				return attendancePeriod, nil
			},
			ClosePeriodFunc: func(id uint, aid uint) error { return nil },
		},
		&mocks.MockUserController{
			GetAllEmployeeFunc: func() ([]models.User, error) {
				return []models.User{}, errors.New("error when get all employee")
			},
		},
		&mocks.MockOvertimeController{},
		&mocks.MockReimbursementController{},
		&mocks.MockPayslipController{},
	)

	req := dto.RunPayrollRequest{AttendancePeriodID: 1}
	err := mockService.RunPayroll(req, 99)
	if err == nil || err.Error() != "error when get all employee" {
		t.Fatalf("expected get all employee error, got: %v", err)
	}
}
