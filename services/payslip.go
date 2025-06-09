package services

import (
	"errors"
	"payslip-system/dto"
	"payslip-system/interfaces"
	"payslip-system/models"
	"time"
)

type PayslipService struct {
	attendanceCtrl    interfaces.AttendanceControllerInterface
	overtimeCtrl      interfaces.OvertimeControllerInterface
	reimbursementCtrl interfaces.ReimbursementControllerInterface
	payslipCtrl       interfaces.PayslipControllerInterface
	payrollCtrl       interfaces.PayrollControllerInterface
}

func NewPayslipService(attendanceCtrl interfaces.AttendanceControllerInterface, overtimeCtrl interfaces.OvertimeControllerInterface, reimbursementCtrl interfaces.ReimbursementControllerInterface, payslipCtrl interfaces.PayslipControllerInterface, payrollCtrl interfaces.PayrollControllerInterface) *PayslipService {
	return &PayslipService{attendanceCtrl: attendanceCtrl, overtimeCtrl: overtimeCtrl, reimbursementCtrl: reimbursementCtrl, payslipCtrl: payslipCtrl, payrollCtrl: payrollCtrl}
}

func (s *PayslipService) GeneratePayslip(userID uint, payrollID uint) (*dto.PayslipResponse, error) {
	payroll, err := s.payrollCtrl.GetPayroll(payrollID)
	if err != nil {
		return nil, errors.New("failed to get payroll")
	}

	attendancePeriod, err := s.attendanceCtrl.GetAttendancePeriodByID(payroll.AttendancePeriodID)
	if err != nil {
		return nil, errors.New("failed to get attendance period")
	}

	payslip, err := s.payslipCtrl.GetPayslipByPayrollID(userID, payrollID)
	if err != nil {
		return nil, errors.New("failed to get payslip")
	}

	attendances, _ := s.attendanceCtrl.GetAttendanceByUserAndDateBetween(userID, attendancePeriod.StartDate, attendancePeriod.EndDate)

	reims, err := s.reimbursementCtrl.GetReimbursements(userID, attendancePeriod.StartDate, attendancePeriod.EndDate)
	if err != nil {
		return nil, err
	}

	overtimes, err := s.overtimeCtrl.GetOvertimeGroupedByDate(userID, attendancePeriod.StartDate, attendancePeriod.EndDate)
	if err != nil {
		return nil, err
	}

	var overtimeResponse []struct {
		Date   time.Time
		Hours  float64
		Rate   float64
		Amount float64
	}

	for _, ovt := range overtimes {
		overtimeResponse = append(overtimeResponse, struct {
			Date   time.Time
			Hours  float64
			Rate   float64
			Amount float64
		}{
			Date:   ovt.Date,
			Hours:  ovt.Hours,
			Rate:   payslip.RateSalaryPerHour,
			Amount: ovt.Hours * payslip.RateSalaryPerHour,
		})
	}

	return &dto.PayslipResponse{
		PeriodStart: attendancePeriod.StartDate,
		PeriodEnd:   attendancePeriod.EndDate,
		Attendance: struct {
			TotalDays int
			Rate      float64
			Amount    float64
		}{
			TotalDays: len(attendances),
			Rate:      payslip.RateSalaryPerDay,
			Amount:    float64(len(attendances)) * payslip.RateSalaryPerDay,
		},
		Overtime:             overtimeResponse,
		Reimbursements:       reims,
		FullAttendanceSalary: 22 * payslip.RateSalaryPerDay,
		TotalTakeHomePay:     payslip.TotalTakeHome,
	}, nil
}

func (s *PayslipService) GetPayslipSummary(payrollID uint) (*dto.PayslipSummary, error) {
	var payslips []models.Payslip
	payslips, err := s.payslipCtrl.GetPayslips(payrollID)
	if err != nil {
		return nil, err
	}

	var summary dto.PayslipSummary
	for _, payslip := range payslips {
		summary.Employees = append(summary.Employees, struct {
			UserID   uint
			FullName string
			TotalPay float64
		}{
			UserID:   payslip.UserID,
			FullName: payslip.User.Username,
			TotalPay: payslip.TotalTakeHome,
		})
		summary.GrandTotal += payslip.TotalTakeHome
	}

	return &summary, nil
}
