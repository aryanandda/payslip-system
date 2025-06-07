package services

import (
	"errors"
	"time"
	"payslip-system/controllers"
	"payslip-system/models"
)

type AttendanceService struct {
	attendanceCtrl *controllers.AttendanceController
}

func NewAttendanceService(attCtrl *controllers.AttendanceController) *AttendanceService {
	return &AttendanceService{
		attendanceCtrl: attCtrl,
	}
}

// Main method - receives userID, processes attendance logic, calls controller for DB ops
func (s *AttendanceService) SubmitAttendance(userID uint) error {
	now := time.Now()

	// Business logic: no weekends
	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		return errors.New("attendance not allowed on weekends")
	}

	// Check DB via controller
	existing, err := s.attendanceCtrl.GetAttendanceByUserAndDate(userID, now)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("attendance already submitted for today")
	}

	// Create record
	attendance := &models.Attendance{
		UserID: userID,
		Date:   now,
	}

	// Persist via controller
	if err := s.attendanceCtrl.CreateAttendance(attendance); err != nil {
		return err
	}

	return nil
}
