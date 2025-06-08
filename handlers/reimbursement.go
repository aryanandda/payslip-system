package handlers

import (
	"net/http"

	"payslip-system/constants"
	"payslip-system/dto"
	"payslip-system/services"

	"github.com/gin-gonic/gin"
)

// SubmitReimbursementHandler godoc
// @Summary      Submit reimbursement
// @Description  Submit a reimbursement request for the logged-in user
// @Tags         reimbursement
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.ReimbursementRequest true "Reimbursement details"
// @Success      200 {object} map[string]string "message: Reimbursement submitted successfully"
// @Failure      400 {object} map[string]string "Invalid request or submission error"
// @Failure      401 {object} map[string]string "Unauthorized"
// @Router       /api/reimbursement [post]
func SubmitReimbursementHandler(svc *services.ReimbursementService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUnauthorized})
			return
		}

		var req dto.ReimbursementRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequest})
			return
		}

		err := svc.SubmitReimbursement(userID.(uint), req.Amount, req.Description)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Reimbursement submitted successfully"})
	}
}
