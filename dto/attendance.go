package dto

type AttendancePeriodRequest struct {
	StartDate string `json:"start_date" binding:"required"` // format: YYYY-MM-DD
	EndDate   string `json:"end_date" binding:"required"`   // format: YYYY-MM-DD
}
