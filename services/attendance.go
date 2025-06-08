package services

import (
	"errors"
	"payslip-system/constants"
	"payslip-system/controllers"
	"payslip-system/models"
	"time"
)

type AttendanceService struct {
	attendanceCtrl *controllers.AttendanceController
}

func NewAttendanceService(attCtrl *controllers.AttendanceController) *AttendanceService {
	return &AttendanceService{
		attendanceCtrl: attCtrl,
	}
}

func (s *AttendanceService) ValidatePeriod(startStr, endStr string, userID uint) (*models.AttendancePeriod, error) {
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

	overlap, err := s.attendanceCtrl.CheckOverlap(startDate, endDate)
	if err != nil {
		return nil, err
	}
	if overlap {
		return nil, errors.New("overlapping attendance period exists")
	}

	return &models.AttendancePeriod{
		StartDate: startDate,
		EndDate:   endDate,
		CreatedBy: &userID,
	}, nil
}

func (s *AttendanceService) CreatePeriod(payload *models.AttendancePeriod) error {
	if err := s.attendanceCtrl.CreateAttendancePeriod(payload); err != nil {
		return errors.New("failed to create attendance period")
	}

	return nil
}

func (s *AttendanceService) SubmitAttendance(userID uint) error {
	now := time.Now()

	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		return errors.New(constants.ErrWeekendAttendance)
	}

	existing, err := s.attendanceCtrl.GetAttendanceByUserAndDate(userID, now)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New(constants.ErrAttendanceAlreadyExists)
	}

	attendance := &models.Attendance{
		UserID: userID,
		Date:   now,
	}

	if err := s.attendanceCtrl.CreateAttendance(attendance); err != nil {
		return err
	}

	return nil
}
