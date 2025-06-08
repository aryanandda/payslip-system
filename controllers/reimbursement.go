package controllers

import (
	"payslip-system/models"
	"time"

	"gorm.io/gorm"
)

type ReimbursementController struct {
	DB *gorm.DB
}

func NewReimbursementController(db *gorm.DB) *ReimbursementController {
	return &ReimbursementController{DB: db}
}

func (c *ReimbursementController) CreateReimbursement(userID uint, amount float64, description string) error {
	reimbursement := models.Reimbursement{
		UserID:      userID,
		Amount:      amount,
		Description: description,
	}
	return c.DB.Create(&reimbursement).Error
}

func (c *ReimbursementController) GetReimbursements(userID uint, start, end time.Time) ([]models.Reimbursement, error) {
	var list []models.Reimbursement
	err := c.DB.Where("user_id = ? AND created_at BETWEEN ? AND ?", userID, start, end).
		Find(&list).Error
	return list, err
}
