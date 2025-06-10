package handlers

import (
	"net/http"
	"payslip-system/constants"
	"payslip-system/services"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Auth
// @Summary      Login
// @Description  Authenticates a user and returns a JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "Login credentials"
// @Success      200 {object} map[string]string "token"
// @Failure      400 {object} map[string]string "Invalid request"
// @Failure      401 {object} map[string]string "Unauthorized"
// @Router       /api/login [post]
func LoginHandler(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRequest})
			return
		}

		token, err := authService.Login(req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func LogoutHandler(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		err := authService.Logout(tokenString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to logout"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
	}
}
