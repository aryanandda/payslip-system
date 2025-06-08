package routes

import (
	"payslip-system/controllers"
	"payslip-system/handlers"
	"payslip-system/middlewares"
	"payslip-system/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB, jwtKey []byte) {
	userCtrl := controllers.NewUserController(db)
	attendanceCtrl := controllers.NewAttendanceController(db)
	payrollCtrl := controllers.NewPayrollController(db)
	overtimeCtrl := controllers.NewOvertimeController(db)
	reimbursementCtrl := controllers.NewReimbursementController(db)
	payslipCtrl := controllers.NewPayslipController(db)

	authSvc := services.NewAuthService(userCtrl, jwtKey)
	attendanceSvc := services.NewAttendanceService(attendanceCtrl)
	payrollSvc := services.NewPayrollService(payrollCtrl, attendanceCtrl, userCtrl, overtimeCtrl, reimbursementCtrl, payslipCtrl)
	overtimeSvc := services.NewOvertimeService(overtimeCtrl)
	reimbursementSvc := services.NewReimbursementService(reimbursementCtrl)
	payslipSvc := services.NewPayslipService(attendanceCtrl, overtimeCtrl, reimbursementCtrl, payslipCtrl, payrollCtrl)

	api := router.Group("/api")

	api.POST("/login", handlers.LoginHandler(authSvc))
	api.POST("/logout", handlers.LogoutHandler(authSvc))

	protected := api.Group("/")
	protected.Use(middlewares.AuthMiddleware(jwtKey))

	protected.POST("/attendance-period", middlewares.RoleMiddleware(), handlers.AddAttendancePeriodHandler(attendanceSvc))
	protected.POST("/payroll/run", middlewares.RoleMiddleware(), handlers.RunPayrollHandler(payrollSvc))
	protected.GET("/payslip-summary/:id", middlewares.RoleMiddleware(), handlers.GetPayslipSummaryHandler(payslipSvc))

	protected.POST("/attendance", handlers.SubmitAttendanceHandler(attendanceSvc))
	protected.POST("/overtime", handlers.SubmitOvertimeHandler(overtimeSvc))
	protected.POST("/reimbursement", handlers.SubmitReimbursementHandler(reimbursementSvc))
	protected.GET("/payslip/:id", handlers.GeneratePayslipHandler(payslipSvc))

}
