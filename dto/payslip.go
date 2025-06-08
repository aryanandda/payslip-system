package dto

import (
	"time"

	"payslip-system/models"
)

type PayslipResponse struct {
	PeriodStart time.Time `json:"period_start"`
	PeriodEnd   time.Time `json:"period_end"`
	Attendance  struct {
		TotalDays int
		Rate      float64
		Amount    float64
	}
	Overtime []struct {
		Date   time.Time
		Hours  float64
		Rate   float64
		Amount float64
	}
	Reimbursements       []models.Reimbursement
	FullAttendanceSalary float64
	TotalTakeHomePay     float64
}

type PayslipSummary struct {
	Employees []struct {
		UserID   uint
		FullName string
		TotalPay float64
	}
	GrandTotal float64
}
