package mocks

import (
	"time"

	"payslip-system/models"
)

type MockAttendanceController struct {
	CheckOverlapFunc                      func(time.Time, time.Time) (bool, error)
	CreateAttendancePeriodFunc            func(*models.AttendancePeriod) error
	GetAttendanceByUserAndDateFunc        func(uint, time.Time) (*models.Attendance, error)
	CreateAttendanceFunc                  func(*models.Attendance) error
	GetAttendancePeriodByIDFunc           func(id uint) (*models.AttendancePeriod, error)
	CountAttendanceFunc                   func(userID uint, start, end time.Time) (int64, error)
	ClosePeriodFunc                       func(id uint, userID uint) error
	GetAttendanceByUserAndDateBetweenFunc func(userID uint, startDate, endDate time.Time) ([]models.Attendance, error)
}

func (m *MockAttendanceController) CheckOverlap(start, end time.Time) (bool, error) {
	return m.CheckOverlapFunc(start, end)
}

func (m *MockAttendanceController) CreateAttendancePeriod(period *models.AttendancePeriod) error {
	return m.CreateAttendancePeriodFunc(period)
}

func (m *MockAttendanceController) GetAttendanceByUserAndDate(userID uint, date time.Time) (*models.Attendance, error) {
	return m.GetAttendanceByUserAndDateFunc(userID, date)
}

func (m *MockAttendanceController) CreateAttendance(att *models.Attendance) error {
	return m.CreateAttendanceFunc(att)
}

func (m *MockAttendanceController) GetAttendancePeriodByID(id uint) (*models.AttendancePeriod, error) {
	return m.GetAttendancePeriodByIDFunc(id)
}

func (m *MockAttendanceController) CountAttendance(userID uint, start, end time.Time) (int64, error) {
	return m.CountAttendanceFunc(userID, start, end)
}

func (m *MockAttendanceController) ClosePeriod(id uint, userID uint) error {
	return m.ClosePeriodFunc(id, userID)
}

func (m *MockAttendanceController) GetAttendanceByUserAndDateBetween(userID uint, startDate, endDate time.Time) ([]models.Attendance, error) {
	return m.GetAttendanceByUserAndDateBetweenFunc(userID, startDate, endDate)
}
