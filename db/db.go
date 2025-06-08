package db

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"golang.org/x/crypto/bcrypt"

	"payslip-system/models"
)

func ConnectAndMigrate(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get generic DB object: %w", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Connected to the database")

	err = db.AutoMigrate(
		&models.User{},
		&models.AttendancePeriod{},
		&models.Attendance{},
		&models.Payroll{},
		&models.Payslip{},
		&models.AuditLog{},
		&models.Overtime{},
		&models.Reimbursement{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to run auto migration: %w", err)
	}

	log.Println("Auto-migration complete")

	if err := seedInitialData(db); err != nil {
		return nil, fmt.Errorf("failed to seed initial data: %w", err)
	}

	return db, nil
}

func seedInitialData(db *gorm.DB) error {
	var count int64
	db.Model(&models.User{}).Count(&count)

	if count > 0 {
		log.Println("Users already exist, skipping seeding.")
		return nil
	}

	log.Println("Seeding default admin and employees...")

	adminPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	admin := models.User{
		Username:     "admin",
		PasswordHash: string(adminPassword),
		IsAdmin:      true,
		Salary:       0,
	}
	if err := db.Create(&admin).Error; err != nil {
		return fmt.Errorf("failed to create admin user: %w", err)
	}

	for i := 1; i <= 100; i++ {
		username := fmt.Sprintf("employee%03d", i)
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		salary := rand.Intn(5000000) + 5000000 // 5.000.000 - 10.000.000

		emp := models.User{
			Username:     username,
			PasswordHash: string(passwordHash),
			IsAdmin:      false,
			Salary:       salary,
		}

		if err := db.Create(&emp).Error; err != nil {
			return fmt.Errorf("failed to create employee %s: %w", username, err)
		}
	}

	log.Println("Seeding complete")
	return nil
}
