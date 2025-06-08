package handlers

import (
	"net/http"

	"payslip-system/constants"
	"payslip-system/dto"
	"payslip-system/services"

	"github.com/gin-gonic/gin"
)

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
