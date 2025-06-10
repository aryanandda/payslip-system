package handlers

import (
	"net/http"
	"payslip-system/constants"
	"payslip-system/dto"
	"payslip-system/services"

	"github.com/gin-gonic/gin"
)

// RunPayrollHandler
// @Summary      Run payroll
// @Description  Processes payroll for employees for a given period (admin only)
// @Tags         payroll
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.RunPayrollRequest true "Payroll run details"
// @Success      200 {object} map[string]string "message: Payroll processed successfully"
// @Failure      400 {object} map[string]string "Invalid request or processing error"
// @Failure      401 {object} map[string]string "Unauthorized"
// @Router       /api/payroll/run [post]
func RunPayrollHandler(payrollSvc *services.PayrollService) gin.HandlerFunc {
	return func(c *gin.Context) {
		adminID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
			return
		}

		var req dto.RunPayrollRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequest})
			return
		}

		if err := payrollSvc.RunPayroll(req, adminID.(uint)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Payroll processed successfully"})
	}
}
