package services

import (
	"errors"
	"payslip-system/constants"
	"payslip-system/mocks"
	"payslip-system/models"
	"testing"
	"time"
)

func TestValidatePeriod_ValidDates(t *testing.T) {
	mockCtrl := &mocks.MockAttendanceController{
		CheckOverlapFunc: func(start, end time.Time) (bool, error) {
			return false, nil
		},
	}
	service := NewAttendanceService(mockCtrl)

	start := "2023-01-01"
	end := "2023-01-10"
	userID := uint(1)

	period, err := service.ValidatePeriod(start, end, userID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if period.StartDate.Format("2006-01-02") != start {
		t.Errorf("StartDate mismatch")
	}
}

func TestValidatePeriod_InvalidStartDate(t *testing.T) {
	service := NewAttendanceService(nil)

	_, err := service.ValidatePeriod("invalid-date", "2023-01-10", 1)
	if err == nil || err.Error() != "invalid start_date format" {
		t.Errorf("Expected invalid start_date format error, got %v", err)
	}
}

func TestValidatePeriod_InvalidEndDate(t *testing.T) {
	service := NewAttendanceService(nil)

	_, err := service.ValidatePeriod("2023-01-10", "invalid-date", 1)
	if err == nil || err.Error() != "invalid end_date format" {
		t.Errorf("Expected invalid end_date format error, got %v", err)
	}
}

func TestValidatePeriod_StartAfterEnd(t *testing.T) {
	service := NewAttendanceService(nil)
	_, err := service.ValidatePeriod("2023-01-10", "2023-01-01", 1)
	if err == nil || err.Error() != "start_date must be before end_date" {
		t.Errorf("expected start before end error, got: %v", err)
	}
}

func TestValidatePeriod_OverlapExists(t *testing.T) {
	mock := &mocks.MockAttendanceController{
		CheckOverlapFunc: func(start, end time.Time) (bool, error) {
			return true, nil
		},
	}

	service := NewAttendanceService(mock)
	_, err := service.ValidatePeriod("2023-01-01", "2023-01-10", 1)
	if err == nil || err.Error() != "overlapping attendance period exists" {
		t.Errorf("expected overlap error, got: %v", err)
	}
}

func TestValidatePeriod_OverlapCheckError(t *testing.T) {
	mock := &mocks.MockAttendanceController{
		CheckOverlapFunc: func(start, end time.Time) (bool, error) {
			return false, errors.New("database error")
		},
	}

	service := NewAttendanceService(mock)
	_, err := service.ValidatePeriod("2023-01-01", "2023-01-10", 1)
	if err == nil || err.Error() != "database error" {
		t.Errorf("expected database error, got: %v", err)
	}
}

func TestSubmitAttendance_WeekdaySuccess(t *testing.T) {
	mockCtrl := &mocks.MockAttendanceController{
		GetAttendanceByUserAndDateFunc: func(userID uint, date time.Time) (*models.Attendance, error) {
			return nil, nil
		},
		CreateAttendanceFunc: func(att *models.Attendance) error {
			if att.UserID != 1 {
				t.Errorf("unexpected UserID")
			}
			return nil
		},
	}

	service := NewAttendanceService(mockCtrl)

	service.nowFunc = func() time.Time {
		return time.Date(2025, 6, 9, 9, 0, 0, 0, time.UTC)
	}

	err := service.SubmitAttendance(1)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestSubmitAttendance_AlreadyExists(t *testing.T) {
	mockCtrl := &mocks.MockAttendanceController{
		GetAttendanceByUserAndDateFunc: func(userID uint, date time.Time) (*models.Attendance, error) {
			return &models.Attendance{UserID: userID, Date: date}, nil // already exists
		},
	}

	service := NewAttendanceService(mockCtrl)
	service.nowFunc = func() time.Time {
		return time.Date(2023, 6, 5, 10, 0, 0, 0, time.UTC) // Monday
	}

	err := service.SubmitAttendance(1)
	if err == nil || err.Error() != constants.ErrAttendanceAlreadyExists {
		t.Errorf("expected duplicate error, got %v", err)
	}
}

func TestSubmitAttendance_Weekend(t *testing.T) {
	service := NewAttendanceService(nil)
	service.nowFunc = func() time.Time {
		return time.Date(2023, 6, 9, 10, 0, 0, 0, time.UTC).AddDate(0, 0, 1) // Saturday
	}

	err := service.SubmitAttendance(1)
	if err == nil || err.Error() != constants.ErrWeekendAttendance {
		t.Errorf("expected weekend error, got %v", err)
	}
}
