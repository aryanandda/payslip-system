package dto

type RunPayrollRequest struct {
	AttendancePeriodID uint `json:"attendance_period_id" binding:"required"`
}
