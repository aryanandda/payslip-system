package interfaces

import (
	"payslip-system/models"
	"time"
)

type AttendanceControllerInterface interface {
	CheckOverlap(start, end time.Time) (bool, error)
	CreateAttendancePeriod(*models.AttendancePeriod) error
	GetAttendanceByUserAndDate(userID uint, date time.Time) (*models.Attendance, error)
	CreateAttendance(*models.Attendance) error
	GetAttendancePeriodByID(id uint) (*models.AttendancePeriod, error)
	CountAttendance(userID uint, start, end time.Time) (int64, error)
	ClosePeriod(id uint, userID uint) error
	GetAttendanceByUserAndDateBetween(userID uint, startDate, endDate time.Time) ([]models.Attendance, error)
}
