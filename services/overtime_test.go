package services

import (
	"errors"
	"payslip-system/constants"
	"payslip-system/mocks"
	"testing"
	"time"
)

func TestSubmitOvertime_Success(t *testing.T) {
	mock := &mocks.MockOvertimeController{
		GetOvertimesByDateFunc: func(userID uint, date time.Time) ([]float64, error) {
			return []float64{1.0}, nil
		},
		CreateOvertimeFunc: func(userID uint, date time.Time, hours float64) error {
			return nil
		},
	}

	service := NewOvertimeService(mock)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	fakeNow := time.Date(2025, 6, 9, 18, 0, 0, 0, loc)
	service.nowFunc = func() time.Time {
		return fakeNow
	}

	err := service.SubmitOvertime(1, 1.5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestSubmitOvertime_InvalidHours(t *testing.T) {
	service := NewOvertimeService(nil)
	err := service.SubmitOvertime(1, 0)
	if err == nil || err.Error() != constants.ErrOvertimeInvalidHours {
		t.Errorf("expected invalid hours error, got %v", err)
	}
}

func TestSubmitOvertime_TooManyHours(t *testing.T) {
	service := NewOvertimeService(nil)
	err := service.SubmitOvertime(1, 5)
	if err == nil || err.Error() != constants.ErrOvertimeTooManyHours {
		t.Errorf("expected too many hours error, got %v", err)
	}
}

func TestSubmitOvertime_Before5PM(t *testing.T) {
	mock := &mocks.MockOvertimeController{}

	service := NewOvertimeService(mock)

	// Fake time before 5 PM
	loc, _ := time.LoadLocation("Asia/Jakarta")
	service.nowFunc = func() time.Time {
		return time.Date(2025, 6, 9, 16, 0, 0, 0, loc)
	}

	err := service.SubmitOvertime(1, 2)
	if err == nil || err.Error() != constants.ErrOvertimeBefore5PM {
		t.Errorf("expected before 5 PM error, got %v", err)
	}
}

func TestSubmitOvertime_ExceedsDailyLimit(t *testing.T) {
	mock := &mocks.MockOvertimeController{
		GetOvertimesByDateFunc: func(userID uint, date time.Time) ([]float64, error) {
			return []float64{3}, nil
		},
	}

	service := NewOvertimeService(mock)

	// Fake time after 5 PM
	loc, _ := time.LoadLocation("Asia/Jakarta")
	service.nowFunc = func() time.Time {
		return time.Date(2025, 6, 9, 18, 0, 0, 0, loc)
	}

	err := service.SubmitOvertime(1, 1.0)
	if err == nil || err.Error() != constants.ErrOvertimeExceedsLimit {
		t.Errorf("expected exceeds limit error, got %v", err)
	}
}

func TestSubmitOvertime_DBError(t *testing.T) {
	mock := &mocks.MockOvertimeController{
		GetOvertimesByDateFunc: func(userID uint, date time.Time) ([]float64, error) {
			return nil, errors.New("db error")
		},
	}

	service := NewOvertimeService(mock)

	// Fake time after 5 PM
	loc, _ := time.LoadLocation("Asia/Jakarta")
	service.nowFunc = func() time.Time {
		return time.Date(2025, 6, 9, 18, 0, 0, 0, loc)
	}

	err := service.SubmitOvertime(1, 1.0)
	if err == nil || err.Error() != "db error" {
		t.Errorf("expected db error, got %v", err)
	}
}
