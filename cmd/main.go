package main

import (
	"log"

	"payslip-system/config"
	"payslip-system/db"
	"payslip-system/middlewares"
	"payslip-system/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment config
	cfg := config.Load()

	// Connect to DB and run auto-migrations
	database, err := db.ConnectAndMigrate(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// Initialize Gin engine
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middlewares.RequestLogger())
	r.Use(middlewares.CaptureIP())

	// Register routes (injecting DB + JWT secret)
	routes.RegisterRoutes(r, database, []byte(cfg.JWTSecret))

	// Start server
	log.Printf("Server running on http://localhost:%s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
