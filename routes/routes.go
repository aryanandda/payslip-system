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
	// Controllers
	userCtrl := controllers.NewUserController(db)
	attendanceCtrl := controllers.NewAttendanceController(db)
	payrollCtrl := controllers.NewPayrollController(db)

	// Services
	authSvc := services.NewAuthService(userCtrl, jwtKey)
	attendanceSvc := services.NewAttendanceService(attendanceCtrl)
	payrollSvc := services.NewPayrollService(payrollCtrl)

	api := router.Group("/api")

	api.POST("/login", handlers.LoginHandler(authSvc))
	api.POST("/logout", handlers.LogoutHandler(authSvc))

	protected := api.Group("/")
	protected.Use(middlewares.AuthMiddleware(jwtKey))

	protected.POST("/attendance", handlers.SubmitAttendanceHandler(attendanceSvc, attendanceCtrl))
	protected.POST("/payroll-period", middlewares.RoleMiddleware(), handlers.AddPayrollPeriodHandler(payrollSvc, payrollCtrl))

}
