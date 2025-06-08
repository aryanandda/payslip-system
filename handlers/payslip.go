package handlers

import (
	"net/http"
	"strconv"

	"payslip-system/constants"
	"payslip-system/services"

	"github.com/gin-gonic/gin"
)

func GeneratePayslipHandler(payslipService *services.PayslipService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
			return
		}

		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payroll ID"})
			return
		}

		payslip, err := payslipService.GeneratePayslip(userID.(uint), uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, payslip)
	}
}

func GetPayslipSummaryHandler(payslipSvc *services.PayslipService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payroll ID"})
			return
		}

		summary, err := payslipSvc.GetPayslipSummary(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, summary)
	}
}
