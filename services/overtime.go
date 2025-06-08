package services

import (
	"errors"
	"fmt"
	"log"
	"payslip-system/constants"
	"payslip-system/controllers"
	"time"
)

type OvertimeService struct {
	ctrl *controllers.OvertimeController
}

func NewOvertimeService(ctrl *controllers.OvertimeController) *OvertimeService {
	return &OvertimeService{ctrl: ctrl}
}

func (s *OvertimeService) SubmitOvertime(userID uint, hours float64) error {
	if hours <= 0 {
		return errors.New(constants.ErrOvertimeInvalidHours)
	}
	if hours > 3 {
		return errors.New(constants.ErrOvertimeTooManyHours)
	}

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Error loading location:", err)
	}

	now := time.Now().In(loc)
	today := time.Now().Truncate(24 * time.Hour)

	cutoff := time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, now.Location())

	log.Printf("now: %v", now)
	log.Printf("cutoff: %v", cutoff)

	if now.Before(cutoff) {
		return errors.New(constants.ErrOvertimeBefore5PM)
	}

	totalHours, err := s.ctrl.GetOvertimesByDate(userID, today)
	if err != nil {
		return err
	}

	var total float64
	for _, hour := range totalHours {
		total += hour
	}

	if total+hours > 3 {
		return errors.New(constants.ErrOvertimeExceedsLimit)
	}

	return s.ctrl.CreateOvertime(userID, today, hours)
}
