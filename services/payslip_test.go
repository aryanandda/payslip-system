package services

import (
	"errors"
	"payslip-system/mocks"
	"payslip-system/models"
	"testing"
	"time"
)

func TestGeneratePayslip_Success(t *testing.T) {
	startDate := time.Date(2025, 5, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, 5, 31, 0, 0, 0, 0, time.UTC)

	service := NewPayslipService(
		&mocks.MockAttendanceController{
			GetAttendancePeriodByIDFunc: func(id uint) (*models.AttendancePeriod, error) {
				return &models.AttendancePeriod{StartDate: startDate, EndDate: endDate}, nil
			},
			GetAttendanceByUserAndDateBetweenFunc: func(uid uint, start, end time.Time) ([]models.Attendance, error) {
				return []models.Attendance{{UserID: uid}}, nil
			},
		},
		&mocks.MockOvertimeController{
			GetOvertimeGroupedByDateFunc: func(uid uint, start, end time.Time) ([]struct {
				Date  time.Time
				Hours float64
			}, error) {
				return []struct {
					Date  time.Time
					Hours float64
				}{
					{Date: startDate, Hours: 2.0},
				}, nil
			},
		},
		&mocks.MockReimbursementController{
			GetReimbursementsFunc: func(uid uint, start, end time.Time) ([]models.Reimbursement, error) {
				return []models.Reimbursement{{Amount: 100000}}, nil
			},
		},
		&mocks.MockPayslipController{
			GetPayslipByPayrollIDFunc: func(uid, pid uint) (*models.Payslip, error) {
				return &models.Payslip{
					RateSalaryPerHour:  50000,
					RateSalaryPerDay:   400000,
					TotalTakeHome:      900000,
					OvertimeHours:      2,
					OvertimePay:        100000,
					ReimbursementTotal: 100000,
				}, nil
			},
		},
		&mocks.MockPayrollController{
			GetPayrollFunc: func(pid uint) (*models.Payroll, error) {
				return &models.Payroll{ID: pid, AttendancePeriodID: 1}, nil
			},
		},
	)

	result, err := service.GeneratePayslip(1, 123)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if result.TotalTakeHomePay != 900000 {
		t.Errorf("Expected total take home pay 900000, got %v", result.TotalTakeHomePay)
	}
}

func TestGeneratePayslip_GetPayrollError(t *testing.T) {
	service := NewPayslipService(
		nil,
		nil,
		nil,
		nil,
		&mocks.MockPayrollController{
			GetPayrollFunc: func(pid uint) (*models.Payroll, error) {
				return nil, errors.New("db error")
			},
		},
	)

	_, err := service.GeneratePayslip(1, 123)
	if err == nil || err.Error() != "failed to get payroll" {
		t.Fatalf("expected 'failed to get payroll' error, got %v", err)
	}
}

func TestGeneratePayslip_GetAttendancePeriodError(t *testing.T) {
	service := NewPayslipService(
		&mocks.MockAttendanceController{
			GetAttendancePeriodByIDFunc: func(id uint) (*models.AttendancePeriod, error) {
				return nil, errors.New("db error")
			},
		},
		nil,
		nil,
		nil,
		&mocks.MockPayrollController{
			GetPayrollFunc: func(pid uint) (*models.Payroll, error) {
				return &models.Payroll{ID: pid, AttendancePeriodID: 1}, nil
			},
		},
	)

	_, err := service.GeneratePayslip(1, 123)
	if err == nil || err.Error() != "failed to get attendance period" {
		t.Fatalf("expected 'failed to get attendance period' error, got %v", err)
	}
}

func TestGeneratePayslip_GetPayslipError(t *testing.T) {
	start := time.Now()
	end := start.AddDate(0, 0, 5)

	service := NewPayslipService(
		&mocks.MockAttendanceController{
			GetAttendancePeriodByIDFunc: func(id uint) (*models.AttendancePeriod, error) {
				return &models.AttendancePeriod{StartDate: start, EndDate: end}, nil
			},
		},
		nil,
		nil,
		&mocks.MockPayslipController{
			GetPayslipByPayrollIDFunc: func(uid, pid uint) (*models.Payslip, error) {
				return nil, errors.New("db error")
			},
		},
		&mocks.MockPayrollController{
			GetPayrollFunc: func(pid uint) (*models.Payroll, error) {
				return &models.Payroll{ID: pid, AttendancePeriodID: 1}, nil
			},
		},
	)

	_, err := service.GeneratePayslip(1, 123)
	if err == nil || err.Error() != "failed to get payslip" {
		t.Fatalf("expected 'failed to get payslip' error, got %v", err)
	}
}

func TestGeneratePayslip_GetReimbursementsError(t *testing.T) {
	start := time.Now()
	end := start.AddDate(0, 0, 5)

	service := NewPayslipService(
		&mocks.MockAttendanceController{
			GetAttendancePeriodByIDFunc: func(id uint) (*models.AttendancePeriod, error) {
				return &models.AttendancePeriod{StartDate: start, EndDate: end}, nil
			},
			GetAttendanceByUserAndDateBetweenFunc: func(uid uint, start, end time.Time) ([]models.Attendance, error) {
				return []models.Attendance{{UserID: uid}}, nil
			},
		},
		nil,
		&mocks.MockReimbursementController{
			GetReimbursementsFunc: func(uid uint, start, end time.Time) ([]models.Reimbursement, error) {
				return nil, errors.New("reimbursement error")
			},
		},
		&mocks.MockPayslipController{
			GetPayslipByPayrollIDFunc: func(uid, pid uint) (*models.Payslip, error) {
				return &models.Payslip{}, nil
			},
		},
		&mocks.MockPayrollController{
			GetPayrollFunc: func(pid uint) (*models.Payroll, error) {
				return &models.Payroll{ID: pid, AttendancePeriodID: 1}, nil
			},
		},
	)

	_, err := service.GeneratePayslip(1, 123)
	if err == nil || err.Error() != "reimbursement error" {
		t.Fatalf("expected 'reimbursement error', got %v", err)
	}
}

func TestGeneratePayslip_GetOvertimeError(t *testing.T) {
	start := time.Now()
	end := start.AddDate(0, 0, 5)

	service := NewPayslipService(
		&mocks.MockAttendanceController{
			GetAttendancePeriodByIDFunc: func(id uint) (*models.AttendancePeriod, error) {
				return &models.AttendancePeriod{StartDate: start, EndDate: end}, nil
			},
			GetAttendanceByUserAndDateBetweenFunc: func(uid uint, start, end time.Time) ([]models.Attendance, error) {
				return []models.Attendance{{UserID: uid}}, nil
			},
		},
		&mocks.MockOvertimeController{
			GetOvertimeGroupedByDateFunc: func(uid uint, start, end time.Time) ([]struct {
				Date  time.Time
				Hours float64
			}, error) {
				return nil, errors.New("overtime error")
			},
		},
		&mocks.MockReimbursementController{
			GetReimbursementsFunc: func(uid uint, start, end time.Time) ([]models.Reimbursement, error) {
				return []models.Reimbursement{}, nil
			},
		},
		&mocks.MockPayslipController{
			GetPayslipByPayrollIDFunc: func(uid, pid uint) (*models.Payslip, error) {
				return &models.Payslip{}, nil
			},
		},
		&mocks.MockPayrollController{
			GetPayrollFunc: func(pid uint) (*models.Payroll, error) {
				return &models.Payroll{ID: pid, AttendancePeriodID: 1}, nil
			},
		},
	)

	_, err := service.GeneratePayslip(1, 123)
	if err == nil || err.Error() != "overtime error" {
		t.Fatalf("expected 'overtime error', got %v", err)
	}
}

func TestGetPayslipSummary_Success(t *testing.T) {
	mockCtrl := &mocks.MockPayslipController{
		GetPayslipsFunc: func(payrollID uint) ([]models.Payslip, error) {
			return []models.Payslip{
				{
					UserID:        1,
					TotalTakeHome: 3000000,
					User:          models.User{Username: "Alice"},
				},
				{
					UserID:        2,
					TotalTakeHome: 2500000,
					User:          models.User{Username: "Bob"},
				},
			}, nil
		},
	}

	service := NewPayslipService(nil, nil, nil, mockCtrl, nil)

	summary, err := service.GetPayslipSummary(123)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(summary.Employees) != 2 {
		t.Errorf("expected 2 employees, got %d", len(summary.Employees))
	}

	if summary.GrandTotal != 5500000 {
		t.Errorf("expected GrandTotal 5500000, got %f", summary.GrandTotal)
	}
}

func TestGetPayslipSummary_GetPayslipsError(t *testing.T) {
	mockCtrl := &mocks.MockPayslipController{
		GetPayslipsFunc: func(payrollID uint) ([]models.Payslip, error) {
			return nil, errors.New("db error")
		},
	}

	service := NewPayslipService(nil, nil, nil, mockCtrl, nil)

	_, err := service.GetPayslipSummary(123)
	if err == nil || err.Error() != "db error" {
		t.Fatalf("expected 'db error', got %v", err)
	}
}
