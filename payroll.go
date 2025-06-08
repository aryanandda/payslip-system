package handlers

import (
	"net/http"

	"payslip-system/constants"
	"payslip-system/controllers"
	"payslip-system/dto"
	"payslip-system/services"

	"github.com/gin-gonic/gin"
)

func AddPayrollPeriodHandler(service *services.PayrollService, controller *controllers.PayrollController) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.AttendancePeriodRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequest})
			return
		}

		userID := c.GetUint("user_id")

		period, err := service.ValidatePeriod(req.StartDate, req.EndDate, userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := controller.CreatePayrollPeriod(period); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create payroll period"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Payroll period created"})
	}
}
