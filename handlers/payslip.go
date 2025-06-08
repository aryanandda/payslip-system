package handlers

import (
	"net/http"
	"strconv"

	"payslip-system/constants"
	"payslip-system/services"

	"github.com/gin-gonic/gin"
)

// GeneratePayslipHandler godoc
// @Summary      Generate payslip
// @Description  Generates a detailed payslip for the logged-in user based on a payroll ID
// @Tags         payslip
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Payroll ID"
// @Success      200 {object} models.Payslip "Generated payslip data"
// @Failure      400 {object} map[string]string "Invalid request or payroll ID"
// @Failure      401 {object} map[string]string "Unauthorized"
// @Failure      500 {object} map[string]string "Internal server error"
// @Router       /api/payslip/{id} [get]
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

// GetPayslipSummaryHandler godoc
// @Summary      Get payslip summary
// @Description  Retrieves a summary of a payslip by payroll ID (for admin or report access)
// @Tags         payslip
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Payroll ID"
// @Success      200 {object} models.PayslipSummary "Payslip summary data"
// @Failure      400 {object} map[string]string "Invalid payroll ID"
// @Failure      500 {object} map[string]string "Internal server error"
// @Router       /api/payslip-summary/{id} [get]
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
