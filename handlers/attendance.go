package handlers

import (
	"net/http"

	"payslip-system/controllers"
	"payslip-system/services"

	"payslip-system/utils"

	"github.com/gin-gonic/gin"
)

func SubmitAttendanceHandler(service *services.AttendanceService, controller *controllers.AttendanceController) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		ip := c.GetString("ip")
		requestID := c.GetString("request_id")

		if err := service.SubmitAttendance(userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		utils.LogAudit(controller.DB, "attendance", "CREATE", 0, userID, ip, requestID)

		c.JSON(http.StatusOK, gin.H{"message": "Attendance submitted"})
	}
}
