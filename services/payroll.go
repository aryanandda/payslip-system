package services

import (
	"errors"
	"payslip-system/dto"
	"payslip-system/interfaces"
	"payslip-system/models"
)

type PayrollService struct {
	payrollCtrl       interfaces.PayrollControllerInterface
	attendanceCtrl    interfaces.AttendanceControllerInterface
	userCtrl          interfaces.UserControllerInterface
	overtimeCtrl      interfaces.OvertimeControllerInterface
	reimbursementCtrl interfaces.ReimbursementControllerInterface
	payslipCtrl       interfaces.PayslipControllerInterface
}

func NewPayrollService(
	payrollCtrl interfaces.PayrollControllerInterface,
	attendanceCtrl interfaces.AttendanceControllerInterface,
	userCtrl interfaces.UserControllerInterface,
	overtimeCtrl interfaces.OvertimeControllerInterface,
	reimbursementCtrl interfaces.ReimbursementControllerInterface,
	payslipCtrl interfaces.PayslipControllerInterface,
) *PayrollService {
	return &PayrollService{
		payrollCtrl:       payrollCtrl,
		attendanceCtrl:    attendanceCtrl,
		userCtrl:          userCtrl,
		overtimeCtrl:      overtimeCtrl,
		reimbursementCtrl: reimbursementCtrl,
		payslipCtrl:       payslipCtrl,
	}
}

func (s *PayrollService) RunPayroll(req dto.RunPayrollRequest, adminID uint) error {
	attendancePeriod, err := s.attendanceCtrl.GetAttendancePeriodByID(req.AttendancePeriodID)
	if err != nil {
		return err
	}
	if attendancePeriod.IsClosed {
		return errors.New("payroll already processed for this period")
	}

	err = s.attendanceCtrl.ClosePeriod(attendancePeriod.ID, adminID)
	if err != nil {
		return errors.New("error when close period")
	}

	payrollID, err := s.payrollCtrl.CreatePayroll(attendancePeriod.ID, adminID)
	if err != nil {
		return errors.New("error when create payroll")
	}

	employees, err := s.userCtrl.GetAllEmployee()
	if err != nil {
		return errors.New("error when get all employee")
	}

	for _, employee := range employees {
		attendances, _ := s.attendanceCtrl.GetAttendanceByUserAndDateBetween(employee.ID, attendancePeriod.StartDate, attendancePeriod.EndDate)
		overtimeTotal, _ := s.overtimeCtrl.GetOvertimeTotal(employee.ID, attendancePeriod.StartDate, attendancePeriod.EndDate)
		reimbursements, _ := s.reimbursementCtrl.GetReimbursements(employee.ID, attendancePeriod.StartDate, attendancePeriod.EndDate)

		totalReimbursement := float64(0)
		for _, reimbursement := range reimbursements {
			totalReimbursement += reimbursement.Amount
		}

		salaryPerDay := float64(employee.Salary) / 22
		salaryPerHour := salaryPerDay / 8
		overtimePay := overtimeTotal * salaryPerHour

		err = s.payslipCtrl.CreatePayslip(models.Payslip{
			UserID:             employee.ID,
			PayrollID:          payrollID,
			OvertimeHours:      overtimeTotal,
			OvertimePay:        overtimePay,
			ReimbursementTotal: totalReimbursement,
			TotalTakeHome:      (float64(len(attendances)) * salaryPerDay) + overtimePay + totalReimbursement,
			PresentDays:        len(attendances),
			CreatedBy:          &adminID,
			RateSalaryPerDay:   salaryPerDay,
			RateSalaryPerHour:  salaryPerHour,
		})
	}

	return nil
}
