package handlers

import (
	"net/http"

	"payslip-system/constants"
	"payslip-system/dto"
	"payslip-system/services"

	"github.com/gin-gonic/gin"
)

// SubmitOvertimeHandler godoc
// @Summary      Submit overtime
// @Description  Submit overtime hours for the logged-in user
// @Tags         overtime
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.OvertimeRequest true "Overtime hours"
// @Success      200 {object} map[string]string "message: Overtime submitted successfully"
// @Failure      400 {object} map[string]string "Invalid request or submission failed"
// @Failure      401 {object} map[string]string "Unauthorized"
// @Router       /api/overtime [post]
func SubmitOvertimeHandler(overtimeSvc *services.OvertimeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
			return
		}

		var req dto.OvertimeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequest})
			return
		}

		err := overtimeSvc.SubmitOvertime(userID.(uint), req.Hours)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Overtime submitted successfully"})
	}
}
