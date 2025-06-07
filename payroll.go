package handlers

import (
	"net/http"

	"payslip-system/controllers"
	"payslip-system/requests"
	"payslip-system/services"

	"github.com/gin-gonic/gin"
)

func AddPayrollPeriodHandler(service *services.PayrollService, controller *controllers.PayrollController) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requests.PayrollPeriodRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		userID := c.GetUint("user_id")

		// Validate and prepare model
		period, err := service.ValidateAndCreatePeriod(req.StartDate, req.EndDate, userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Save to DB
		if err := controller.CreatePayrollPeriod(period); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create payroll period"})
			return
		}

		// Log to audit
		// utils.LogAudit(controller.DB, "payroll_periods", "CREATE", period.ID, userID, ip, requestID)

		c.JSON(http.StatusCreated, gin.H{"message": "Payroll period created"})
	}
}
