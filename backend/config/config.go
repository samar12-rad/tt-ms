package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	user := os.Getenv("PG_USER")
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DATABASE")

	// Get SSL mode from environment variable, default to 'disable' locally
	sslMode := os.Getenv("PG_SSLMODE")
	if sslMode == "" {
		sslMode = "disable" // Local dev default
	}

	// Build DSN dynamically using sslMode
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&connect_timeout=10",
		user, password, host, port, dbname, sslMode,
	)

	// Optional: Supabase sometimes needs this tweak
	if sslMode == "require" {
		dsn += "&prefer_simple_protocol=true"
	}

	log.Printf("Connecting to database at %s with SSL mode: %s", host, sslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
	log.Println("✅ Successfully connected to database")
}
