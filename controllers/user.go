package controllers

import (
	"errors"

	"payslip-system/models"

	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (ctrl *UserController) GetAllEmployee() ([]models.User, error) {
	users := make([]models.User, 0)
	if err := ctrl.DB.Where("is_admin = ?", false).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return users, nil
}

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
