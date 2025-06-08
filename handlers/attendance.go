package handlers

import (
	"net/http"

	"payslip-system/constants"
	"payslip-system/dto"
	"payslip-system/services"

	"github.com/gin-gonic/gin"
)

// AddAttendancePeriodHandler godoc
// @Summary      Add attendance period
// @Description  Creates a new attendance period for the logged-in user
// @Tags         attendance
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.AttendancePeriodRequest true "Attendance Period Dates"
// @Success      201 {object} map[string]string "message: Attendance period created"
// @Failure      400 {object} map[string]string "Invalid request or validation error"
// @Failure      500 {object} map[string]string "Internal server error"
// @Router       /api/attendance-period [post]
func AddAttendancePeriodHandler(service *services.AttendanceService) gin.HandlerFunc {
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

		err = service.CreatePeriod(period)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Attendance period created"})
	}
}

// SubmitAttendanceHandler godoc
// @Summary      Submit attendance
// @Description  Submits the current user's attendance for the active period
// @Tags         attendance
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} map[string]string "message: Attendance submitted"
// @Failure      400 {object} map[string]string "Submission failed"
// @Router       /api/attendance [post]
func SubmitAttendanceHandler(service *services.AttendanceService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")

		if err := service.SubmitAttendance(userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Attendance submitted"})
	}
}
