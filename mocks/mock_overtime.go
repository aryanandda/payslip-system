package mocks

import (
	"time"
)

type MockOvertimeController struct {
	GetOvertimesByDateFunc       func(userID uint, date time.Time) ([]float64, error)
	CreateOvertimeFunc           func(userID uint, date time.Time, hours float64) error
	GetOvertimeTotalFunc         func(userID uint, start, end time.Time) (float64, error)
	GetOvertimeGroupedByDateFunc func(userID uint, start, end time.Time) ([]struct {
		Date  time.Time
		Hours float64
	}, error)
}

func (m *MockOvertimeController) GetOvertimesByDate(userID uint, date time.Time) ([]float64, error) {
	return m.GetOvertimesByDateFunc(userID, date)
}

func (m *MockOvertimeController) CreateOvertime(userID uint, date time.Time, hours float64) error {
	return m.CreateOvertimeFunc(userID, date, hours)
}

func (m *MockOvertimeController) GetOvertimeTotal(userID uint, start, end time.Time) (float64, error) {
	return m.GetOvertimeTotalFunc(userID, start, end)
}

func (m *MockOvertimeController) GetOvertimeGroupedByDate(userID uint, start, end time.Time) ([]struct {
	Date  time.Time
	Hours float64
}, error) {
	return m.GetOvertimeGroupedByDateFunc(userID, start, end)
}