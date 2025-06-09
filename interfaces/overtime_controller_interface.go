package interfaces

import "time"

type OvertimeControllerInterface interface {
	CreateOvertime(userID uint, date time.Time, hours float64) error
	GetOvertimesByDate(userID uint, date time.Time) ([]float64, error)
	GetOvertimeTotal(userID uint, start, end time.Time) (float64, error)
	GetOvertimeGroupedByDate(userID uint, start, end time.Time) ([]struct {
		Date  time.Time
		Hours float64
	}, error)
}
