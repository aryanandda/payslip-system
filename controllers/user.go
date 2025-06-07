package controllers

import (
	"errors"

	"gorm.io/gorm"
	"payslip-system/models"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

// Find user by username
func (ctrl *UserController) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := ctrl.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// For logout or session management, optionally you could have:
// func (ctrl *UserController) UpdateUserToken(userID uint, token string) error { ... }
